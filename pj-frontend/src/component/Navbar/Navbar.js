import React, { Component } from 'react';
import { Button } from "../Button";
import { MenuItems } from  "./MenuItems";
import "./Navbar.css";

//48 mins https://www.youtube.com/watch?v=fL8cFqhTHwA
class Navbar extends Component {
    state = {clicked:false}
handleClick = ()=>{
    this.setState({ clicked: !this.state.clicked })
}

//removed <Button>SignUp</Button> after </ul> because it was causing a second signup button

render(){
    return(
        <nav className="NavbarItems">
        <h1 className="navbar-logo">Pickle Jar<i className="fab fa-react">  
        </i></h1>

        <div className="menu-icon" onClick={this.handleClick}>
        <i className={this.state.clicked ? 'fas fa-times' : 'fas fa-bars'}></i>
            
        </div>
        <ul className={this.state.clicked ? 'nav-menu active' : 
        'nav-menu'}>
            {MenuItems.map((item, index)=>{
                return(
                       <li key={index}>

                       <a className={item.cName} href={item.url}>
                        {item.title}
                            </a>
                       </li>
                )
            })}

         
        </ul>
        </nav>
    )
}


}
export default Navbar
