import { Modal } from "react-responsive-modal";
import "react-responsive-modal/styles.css";
import "./css/surveyModal.css";
import Szme from "../assets/스즈메의 문단속(찐).jpg";
import JujutsuKaisen from "../assets/주술회전.webp";

const ModalExample = (props: any) => {
  function closeModal() {
    props.setOpen(false);
  }
  const data = [
    { img: Szme, title: "스즈메의 문단속" },
    { img: JujutsuKaisen, title: "주술회전" },
    { img: "검정", title: "#000000" },
  ];
  return (
    <>
      <Modal
        open={props.open}
        onClose={closeModal}
        center
        classNames={{ modal: "customModal" }}
      >
        <h2>{props.title}</h2>
        <p>{props.content}</p>
        <form className="formContainer">
          {data.map((img, index) => (
            <div className="Item" key={index}>
              <input
                name="result"
                className="opacity(0)"
                type="checkbox"
                id={`checkbox${index}`}
              />
              <label htmlFor={`checkbox${index}`}>
                <div className={`w(100) h(auto)`}>
                  <div className="font(12) text(center) text w(100)">
                    {img.title}
                  </div>
                  <img
                    src={img.img}
                    className="filter(grayscale(50%)) w(100)"
                  ></img>
                </div>
              </label>
            </div>
          ))}
        </form>
      </Modal>
    </>
  );
};

export default ModalExample;
