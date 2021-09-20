import Logout from "../Logout/Logout";

function Dashboard(props) {

  //{props.user.first_name} use this after welcome if passing user state to dashboard
  console.log(props.user.first_name)
  
  

  return (
    <main className="container center">
      <div id="coolDash">
        <div id="theUser">
        <p>WELCOME, {localStorage.getItem("first_name")} !</p>
        </div>
        <br></br>
        <div id="theOpens">
        <p>The Picklejar Has Been Opened:</p>
        <p>[call to backend] times</p>
        <button>UNLOCK</button>
        <button>LOCK</button>
        </div>
        <br></br>
        <div>
            <p>Packages In The Picklejar Today:</p>
            <p>[call to backend] packages</p>
        </div>
        <br></br>
        <Logout/>
        
      </div>
    </main>
  );
}

export default Dashboard;