import React from 'react';
import pJbp from '../image/pJbp.jpg';


function Home() {

  return (
    <main className="container center">
        <div className="image"
        style={{
          backgroundImage: 'url('+pJbp+')',
          backgroundSize: "cover",
          height: "100vh",
          color: "#f5f5f5"
        }}
      >
      </div>


        <div className="rightside">
          <a href="/SignUp"><button className="coolbutton1"><h1>Sign Up</h1></button></a>
          <br></br>
          <br></br>
          <a href="/Login"><button className="coolbutton2"><h1>Log In</h1></button></a>
        </div>


    </main>
  );
}

export default Home;