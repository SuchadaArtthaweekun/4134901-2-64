import React from "react";



class Wheel extends React.Component {
 
  render(){
    return  <p>I have 4 Wheels.</p>;
  
  }
}

class Car extends React.Component {
    constructor(props){
      super(props);
      this.state = {brand: "Tonz"};
    }
    changeBrand = (msg)=>{
      this.setState({ brand: "Mazda"});
      alert(msg);
    };

    render(){
      return (
        <>
        <Wheel />
        <p>
            My car is <span>{this.state.brand}</span> ,{this.props.age} 
        <br />
            Eye color is {this.props.eye}
        </p> 
        <button onClick={() => this.changeBrand("Hi !!")}>Click Me !!</button>
        </>
      );
    }
  }
  
  export default Car;