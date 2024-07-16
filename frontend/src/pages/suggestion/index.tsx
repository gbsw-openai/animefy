import Chat from "../../components/chat";
import Headers from "../../components/header";
export default function Suggestion() {
  return (
    <>
      <div className="fixed(center) w(80%~500) pack">
        <Chat />
      </div>
      <Headers />
    </>
  );
}
