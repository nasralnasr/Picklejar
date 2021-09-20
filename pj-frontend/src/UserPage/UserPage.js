import { Redirect } from "react-router-dom";
import { Cookies } from "react-cookie";
import Dashboard from "../Dashboard/Dashboard";
import { useEffect, useState } from "react";

function UserPage(props) {
  // Is CreatedAt and UpdatedAt a hard coded string from the server?
  let validUser = false;
  const cookies = new Cookies();
  //const {first_name, last_name, password, email, phone_number} = cookies.get("token");

  const first_name = JSON.stringify(localStorage.getItem("first_name"))
  const last_name = JSON.stringify(localStorage.getItem("last_name"))
  const email = JSON.stringify(localStorage.getItem("email"))
  const password = JSON.stringify(localStorage.getItem("password"))
  const phone_number = JSON.stringify(localStorage.getItem("phone_number"))

  if (first_name && last_name && password && email && phone_number) {
    validUser = true;
    return <Dashboard user={cookies.get("token")} />
  }
  else{
    return <Redirect  to = "/" />
  }
  
  //return validUser  ? <Dashboard user={cookies.get("token")} /> : <Redirect to="/" />;
}

export default UserPage;
