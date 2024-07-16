import { Link } from "react-router-dom";
import Logo from "../assets/DefultLogo.png";

function App() {
  return (
    <>
      <div className="w(100%) vpack">
        <img src={Logo} alt="" className="w(700) mt(80)"></img>
        <Link to={"/suggestion"} className="bold font(20) mt(80)">
          <button className="bold bg(#ffff) c(black) b(3/black) r(40) w(200) h(60) hover:bg(black) hover:c(#ffff) hover:font(18) transition(.5s)">
            시작하기
          </button>
        </Link>
        <div>
          <div className="pack mt(300) gap(50)">
            <div className="w(170) h(150) bg(#000) r(20)"></div>
            <div className="w(500) h(150) bg(#000) r(20) ">
              <div className="c(#fff) vpack text(center) bold font(18) p(30) gap(20)">
                좋아하는 애니를 선택하면 AI 정확도가 올라가요!
                <Link to={"/test"}>
                  <div className="r(50) bg(#fff) w(180) h(50) c(#000) line-height(50) hover:bg(black) hover:c(#ffff) hover:font(17) transition(.5s)">
                    참여 하기
                  </div>
                </Link>
              </div>
            </div>
          </div>
        </div>
        <div>
          <div className="pack mt(20) gap(50)">
            <div className="w(500) h(150) bg(#000) r(20) ">
              <div className="c(#fff) vpack text(center) bold font(18) p(30) gap(20) ">
                좋아하는 애니를 선택하면 AI 정확도가 올라가요!
                <div className="r(50) bg(#fff) w(180) h(50) c(#000) line-height(50) hover:bg(black) hover:c(#ffff) hover:font(17) transition(.5s)">
                  참여 하기
                </div>
              </div>
            </div>
            <div className="w(170) h(150) bg(#000) r(20)"></div>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
