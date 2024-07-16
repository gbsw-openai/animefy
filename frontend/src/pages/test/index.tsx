import { useState } from "react";
import SurveyModal from "../../components/surveyModal";
import Header from "../../components/header";
const Test = () => {
  const [open, setOpen] = useState(false);
  return (
    <>
      <div>
        <button
          className="bold bg(#ffff) c(black) b(3/black) r(40) absolute x(center) y(50%) w(240) h(70) hover:bg(black) hover:c(#ffff) hover:font(18) transition(.5s)    "
          onClick={() => setOpen(!open)}
        >
          설문 조사 시작하기
        </button>
        <SurveyModal
          open={open}
          setOpen={setOpen}
          title="테스트"
          content="이건 테스트 문장입니다."
        />
        <Header />
      </div>
    </>
  );
};

export default Test;
