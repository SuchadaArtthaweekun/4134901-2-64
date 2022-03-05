import React from "react";
import ReactDOM from "react-dom";

import 'bootstrap/dist/css/bootstrap.min.css';
import "./index.css";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import Car from "./components/Car";
import List from "./components/List";
import UI from "./components/Ui";
import Datatable from "./components/Table";
import Hook from "./components/Hook";
import Todo from "./components/Todo";
import Addbar from "./components/Addbar"

const Routing = () => {
  return (
    <Router>
      <Link to="/">Home </Link>
      <Link to="/ui">UI </Link>
      <Link to="/list">List </Link>
      <Link to="/hook">Hook </Link>
      <Link to="/car">Car</Link>
      <Link to="/todo">Todo</Link>
      <Switch>
        <Route exact path="/" component={Datatable} />
        <Route path="/ui" component={UI} />
        <Route path="/list" component={List} />
        <Route path="/hook" component={Hook} />
        <Route path="/car" age="25" eye="black" component={Car} />
        <Route path="/todo" component={Todo} />
      </Switch>
    </Router>
  );
};

ReactDOM.render(
  <>
    <Routing />
    <Addbar />
  </>,
  document.getElementById("root")
);
