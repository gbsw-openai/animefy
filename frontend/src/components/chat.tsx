import ChatBtn from "../assets/발바닥.png";
import "./css/chat.css";

const Chat = () => {
  return (
    <form className="chat-form">
      <textarea className="chat-textarea" placeholder="Type your message..." />
      <button className="chat-button">
        <img src={ChatBtn} alt="Send" width={24} />
      </button>
    </form>
  );
}

export default Chat;