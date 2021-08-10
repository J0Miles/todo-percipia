import React from "react";
import {Header} from "./components/Header";
import {TodoList} from "./components/TodoList";
// import {TodoForm} from "./components/TodoForm";
import "./App.css";

function App() {
  return (
    <div className="App">
      <Header />
      <div>
      </div>
      <TodoList />
    </div>
  );
}

export default App;
