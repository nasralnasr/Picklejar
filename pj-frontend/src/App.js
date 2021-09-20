import React from "react";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "./Home/Home";
import Login from "./Login/Login";
import SignUp from "./SignUp/SignUp";
import Footer from "./Footer/Footer";
import UserPage from "./UserPage/UserPage";
//import PickeHomePageComponent from "./component/Navbar";
import Navbar from "./component/Navbar/Navbar";

function App() {
  return (
    <Router>
      <div className="App">
      <Navbar/>
        <Switch>
          <Route path="/UserPage">
            <UserPage />
          </Route>
          <Route path="/SignUp">
            <SignUp />
          </Route>
          <Route path="/Login">
            <Login />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </div>
      <Footer />
    </Router>
  );
}

export default App;
