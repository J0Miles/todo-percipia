import React from "react";
import { Header } from "./components/Header";
import { TodoList } from "./components/TodoList";
import "./App.css";

function App() {
  return (
    <div className="App">
      <Header />
      <span className="flex justify-between p-6">
        <h2 className="align-left font-medium">Task</h2>
        <h2 className="font-medium align-right">Complete</h2>
      </span>
      <TodoList
        todos={[
          { title: "Do Dishes", created_at: "08/04/2021", completed: false },
          {
            title: "Let the dogs out",
            created_at: "08/04/2021",
            completed: true,
          },
          { title: "Wash Car", created_at: "08/04/2021", completed: false },
          {
            title: "Listen to audiobook",
            description: "The Art of War",
            created_at: "08/04/2021",
            completed: false,
          },
        ]}
      />
    </div>
  );
}

export default App;
