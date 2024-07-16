import { useState, useEffect } from "react";
import ChatBtn from "../assets/발바닥.png";
import "./css/chat.css";
import axios from "axios";

interface Message {
  text: string;
  sender: string;
}

const Chat = () => {
  const [rows, setRows] = useState(1);
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState<Message[]>([]);

  const sendMessage = async (message: string) => {
    try {
      const response = await axios.post('http://localhost:8080/chat/1', { message: message });
      console.log(response.data);

      setMessages([...messages, { text: message, sender: "user" }, { text: response.data.response, sender: "bot" }]);
    } catch (error) {
      console.error("Error sending message:", error);
    }
  }

  const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    const textareaLineHeight = 24;
    const previousRows = event.target.rows;
    event.target.rows = 1;

    const currentRows = Math.floor(
      event.target.scrollHeight / textareaLineHeight
    );

    if (currentRows === previousRows) {
      event.target.rows = currentRows;
    }

    if (currentRows >= 5) {
      event.target.rows = 5;
      event.target.scrollTop = event.target.scrollHeight;
    }

    setRows(currentRows < 5 ? currentRows : 5);
    setMessage(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    sendMessage(message);
    setMessage('');
    setRows(1);
  }

  return (
    <div className="chat-container">
      <div className="messages-container">
        {messages.map((msg, index) => (
          <div key={index} className={`message ${msg.sender}`}>
            {msg.text}
          </div>
        ))}
      </div>
      <form className="chat-form" onSubmit={handleSubmit}>
        <textarea
          className="chat-textarea"
          placeholder="Type your message..."
          rows={rows}
          value={message}
          onChange={handleChange}
        />
        <button className="chat-button" type="submit">
          <img src={ChatBtn} alt="Send" width={24} />
        </button>
      </form>
    </div>
  );
};

export default Chat;