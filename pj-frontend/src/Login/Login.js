import React from "react";
import { useState } from "react";
import "./Login.css";


function Login() {
  const [user, setUser] = useState(null);

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

    fetch("http://localhost:8080/validate-login", {
      method: "POST",
      body: JSON.stringify({ user }),
      headers: { "Content-Type": "application/json" },
    })
      .then((res) => res.json())
      .then((json) => {
        console.log(json);
        setUser(json.user);
      })
      .catch((err) => {
        console.error(err);
      });
  }

  return (
    <main className="Login">
      <div id="coollogin">
        <form onSubmit={handleSubmitUser}>
          <label className="coolUserField">
            <strong>Email:</strong>
          </label>
          <input
            className="coolUserField"
            type="text"
            name="email"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <label className="coolPassField">
            <strong>password:</strong>
          </label>
          <input
            className="coolPassField"
            type="password"
            name="password"
            onChange={handleSetUser}
          ></input>
          <br></br>
          <br></br>
          <input className="loginButtons" type="submit" name="Login" />
        </form>
        <a href="/">
          <button className="loginButtons">Go Back</button>
        </a>
      </div>
    </main>
  );
}

export default Login;