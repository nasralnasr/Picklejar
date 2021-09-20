import React from "react";
import { useState } from "react";
import { useCookies } from "react-cookie";
import { Redirect } from "react-router-dom";
import "./SignUp.css";


function SignUp() {
  const [user, setUser] = useState(null);
  const [redirect, setRedirect] = useState(false);
  const [cookies, setCookie] = useCookies(["token"]);

  function handleSetUser(e) {
    setUser(() => {
      return {
        ...user,
        [e.target.name]: e.target.value,
      };
    });
  }

  function handleSubmitUser(e) {
    //  console.log("email: " + user.email + " password: " + user.password)
    //console.log(user.password)
    e.preventDefault();
    console.log(JSON.stringify({ user }));

    fetch("http://localhost:8080/register", {
      method: "POST",
      body: JSON.stringify({ user }),
      headers: { "Content-Type": "application/json" },
    })
      .then((res) => res.json())
      .then((jwt) => {
        setUser(jwt);
        setCookie("token", jwt, { path: "/" });

         //stores Token to localStorage, if phone number is reused it will return phone number already exists
        localStorage.setItem("token", JSON.stringify(jwt));

        setRedirect(!redirect);
      })
      .catch((err) => {
        console.error(err);
      });

      //Stores user info to localStorage on submit
      localStorage.setItem("first_name", user.first_name);
      localStorage.setItem("last_name", user.last_name);
      localStorage.setItem("email", user.email);
      localStorage.setItem("password", user.password);
      localStorage.setItem("phone_number", user.phone_number);
             
  }

  return (
    <main className="SignUp">
      {redirect && <Redirect to="/UserPage" />}
      <div id="coollogin">
        <form className="coolForm" onSubmit={handleSubmitUser}>
          <label>FirstName:</label>
          <input
            required
            type="text"
            name="first_name"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <label>LastName:</label>
          <input
            required
            type="text"
            name="last_name"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <label>Email:</label>
          <input
            required
            type="text"
            name="email"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <label>
            <strong>Password:</strong>
          </label>
          <input
            required
            type="password"
            name="password"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <label>Phone Number:</label>
          <input
            required
            type="text"
            name="phone_number"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <input type="submit"></input>
        </form>
        <a href="/Home">
          <button className="coolForm">Go Back</button>
        </a>
      </div>
    </main>
  );
}

export default SignUp;
