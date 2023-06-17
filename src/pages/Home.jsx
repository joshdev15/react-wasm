import { useRef, useState } from "react";

const Home = () => {
  const [showErr, setShowErr] = useState(false);
  const [showRes, setShowRes] = useState(false);
  const [result, setResult] = useState();
  const aValue = useRef();
  const bValue = useRef();

  const calculate = async (a, b) => {
    return window.asyncSum(a, b);
  };

  const onSubmit = async (e) => {
    e.preventDefault();

    const aIsNaN = isNaN(aValue.current.value);
    const bIsNaN = isNaN(bValue.current.value);

    if (aIsNaN || bIsNaN) {
      setShowErr(true);
      setTimeout(() => {
        setShowErr(false);
      }, 3000);
      return;
    }

    const aParsed = parseInt(aValue.current.value);
    const bParsed = parseInt(bValue.current.value);

    // setResult(window.sum(aParsed, bParsed))

    calculate(aParsed, bParsed).then((res) => {
      console.log(res);
      setResult(res);
      setShowRes(true);
    });

    console.log("Hola gente");
    // setResult(window.asyncSum(aParsed, bParsed));
  };

  return (
    <div className="home">
      <form onSubmit={(e) => onSubmit(e)}>
        <input ref={aValue} id="valuea" />
        <input ref={bValue} id="valueb" />
        <button type="submit">Calcular</button>
        {showErr && <div>Invalid Value</div>}
        {showRes && <div>{result}</div>}
      </form>
    </div>
  );
};

export default Home;
