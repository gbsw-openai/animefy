import LogoImg from "../assets/로고.png";
import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
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
          className={`absolute w(100%) h(100vh) bg(0,0,0,0.5) opacity(0) ${
            menu ? "fade-in opacity(1)" : "fade-out"
          }`}
          onClick={onToggle}
        >
          <div
            className={`menu-container ${
              menu ? "slide-in" : "slide-out"
            } hbox(top+left) p(10)`}
            onClick={(e) => e.stopPropagation()}
          >
            <button onClick={onToggle} className="c(#fff) font(30) pl(10)">
              {"<"}
            </button>
            <div className="vpack mt(30) c(#fff) absolute(30,30) gap(10) w(200)">
              <Link to={"/evaluation"} className="bb(1) w(100%) bold">
                애니 평가
              </Link>
              <Link to={"/suggestion"} className="bb(1) w(100%) bold">
                애니 추천
              </Link>
            </div>
          </div>
        </div>
      )}
      <div className="hbox space-between w(100%) p(20)">
        <button
          className="b(1) w(40) h(40) bg(rgb(73,73,73)) c(#fff) w(50) h(50) r(8)"
          onClick={onToggle}
        >
          M
        </button>

        <Link to={"http://localhost:5173/"}>
          <img src={LogoImg} alt="" width={160} />
        </Link>
      </div>
    </>
  );
};

export default Header;
