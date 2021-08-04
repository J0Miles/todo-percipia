import React from "react";
import { Todo } from "../models/Todo";
import { TodoItem } from "./TodoItem";

type Props = {
  todos: Todo[];
};

export const TodoList: React.FC<Props> = ({ todos }) => {
  return (
    <ul>
      {todos.map((todo, i) => (
        <li key={i} className="p-4">
          <TodoItem
            title={todo.title}
            description={todo?.description}
            created_at={todo.created_at}
            updated_at={todo?.updated_at}
            completed={todo.completed}
          />
        </li>
      ))}
    </ul>
  );
};
