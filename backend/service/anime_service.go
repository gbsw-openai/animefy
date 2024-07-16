package services

import (
	"context"
	"fmt"
	"strings"

	model "backend/model"
	repo "backend/repository"

	openai "github.com/sashabaranov/go-openai"
)

type ChatService interface {
	GetUserFeed(userID int) ([]model.Anime, error)
	ChatWithBot(userID int, message string) (string, error)
}

type chatService struct {
	animeRepo    repo.AnimeRepository
	openaiClient *openai.Client
}

func NewChatService(animeRepo repo.AnimeRepository, openaiKey string) ChatService {
	client := openai.NewClient(openaiKey)
	return &chatService{
		animeRepo:    animeRepo,
		openaiClient: client,
	}
}

func (cs *chatService) GetUserFeed(userID int) ([]model.Anime, error) {
	return cs.animeRepo.GetAllAnimes()
}

func (cs *chatService) ChatWithBot(userID int, message string) (string, error) {
	animes, err := cs.animeRepo.GetAllAnimes()
	if err != nil {
		return "", fmt.Errorf("failed to get anime data: %v", err)
	}

	animeData := formatAnimeData(animes)
	prompt := fmt.Sprintf(
		`당신은 애니를 추천해주는 친절한 언시스턴스입니다. 사용자의 물음에 친절히 대답해주세요. 자신이 가진 정보를 최대한 활용하여 사용자에게 애니메이션을 추천해 주세요. 당신이 가진 정보는 다음과 같습니다. %s. 당신이 가진 정보, anime_data에는 name, genre, type, episodes, rating, members가 중요합니다. name은 애니의 이름입니다. genre는 애니의 장르입니다. type은 애니의 유형입니다. episodes는 총 방영 수입니다. 마지막 members는 숫자가 적으면 마이너한 애니이고 숫자가 크다면 대중적인 애니입니다. 사용자에게 추천하는 애니를 최대 3개까지만 대답해주고 거짓된 정보는 누설하지 말아주세요. 또한 추천하는 애니의 설명을 자세히 해주세요.`,
		animeData,
	)

	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: prompt},
		{Role: "user", Content: message},
	}

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		Messages:  messages,
		MaxTokens: 300,
	}

	ctx := context.Background()

	resp, err := cs.openaiClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to complete chat request: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func formatAnimeData(animes []model.Anime) string {
	var sb strings.Builder
	count := 0
	for _, anime := range animes {
		sb.WriteString(fmt.Sprintf("Name: %s, Genre: %s, Type: %s, Episodes: %d, Rating: %.2f, Members: %d\n",
			anime.Name, anime.Genre, anime.Anitype, anime.Episodes, anime.Rating, anime.Members))
		count++
		if count >= 100 {
			break
		}
	}
	return sb.String()
}
