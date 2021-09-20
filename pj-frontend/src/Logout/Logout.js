import { Redirect } from "react-router-dom";
import Home from "../Home/Home";
import "./Logout.css";

function Logout(){

    function goBack() {

        localStorage.clear();
        document.location.href="/";

    }


    return (
        <button onClick={goBack}>LOGOUT</button>
    );
}

export default Logout;