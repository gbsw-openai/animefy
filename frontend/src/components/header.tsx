import LogoImg from "../assets/로고.png";
import { useState, useEffect } from "react";
import "./css/header.css";

const Header = () => {
  const [menu, setMenu] = useState(false);
  const [animateMenu, setAnimateMenu] = useState(false);
  const [animating, setAnimating] = useState(false);

  useEffect(() => {
    if (menu) {
      setAnimateMenu(true);
    } else {
      setAnimating(true);
      setTimeout(() => {
        setAnimateMenu(false);
        setAnimating(false);
      }, 400);
    }
  }, [menu]);

  function onToggle() {
    if (!animating) {
      setMenu(!menu);
    }
  }

  return (
    <>
      {(menu || animateMenu) && (
        <div
          className={`absolute w(100%) h(100vh) bg(0,0,0,0.5) opacity(0) ${menu ? "fade-in opacity(1)" : "fade-out"
            }`}
          onClick={onToggle}
        >
          <div
            className={`menu-container ${menu ? "slide-in" : "slide-out"
              } hbox(top+left) p(10)`}
            onClick={(e) => e.stopPropagation()}
          >
            <button onClick={onToggle}>{"<"}</button>
            <div className="vpack mt(30)">
              <div>애니 평가</div>
              <div>애니 추천</div>
              <div>김한결</div>
            </div>
          </div>
        </div>
      )}
      <div className="hbox space-between w(100%) p(20)">
        <button className="b(1) w(40) h(40) bg(#f00)" onClick={onToggle}>
          메뉴
        </button>
        <img src={LogoImg} alt="" width={160} />
      </div>
    </>
  );
};

export default Header;
