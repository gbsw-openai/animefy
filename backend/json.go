package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Anime struct {
	ID       string `json:"anime_id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Type     string `json:"type"`
	Episodes string `json:"episodes"`
	Rating   string `json:"rating"`
	Members  string `json:"members"`
}

func sss() {
	// CSV 파일 열기
	csvFile, err := os.Open("anime.csv")
	if err != nil {
		log.Fatalf("failed to open CSV file: %v", err)
	}
	defer csvFile.Close()

	// CSV 리더 생성
	reader := csv.NewReader(csvFile)

	// 첫 번째 줄은 헤더이므로 건너뜀
	_, err = reader.Read()
	if err != nil {
		log.Fatalf("failed to read CSV header: %v", err)
	}

	var animes []Anime

	// CSV 파일 읽기
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read CSV line: %v", err)
		}

		// CSV 데이터 파싱
		anime := Anime{
			ID:       line[0],
			Name:     line[1],
			Genre:    line[2],
			Type:     line[3],
			Episodes: line[4],
			Rating:   line[5],
			Members:  line[6],
		}

		animes = append(animes, anime)
	}

	// JSON 파일로 저장
	jsonFile, err := os.Create("anime.json")
	if err != nil {
		log.Fatalf("failed to create JSON file: %v", err)
	}
	defer jsonFile.Close()

	// JSON 인코딩
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ") // 들여쓰기 설정

	err = encoder.Encode(animes)
	if err != nil {
		log.Fatalf("failed to encode JSON: %v", err)
	}

	fmt.Println("Conversion successful. JSON file saved as anime.json")
}
