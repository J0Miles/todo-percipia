import React from "react";
import { Todo } from "../models/Todo";

export const TodoItem: React.FC<Todo> = ({
  title,
  description,
  created_at,
  updated_at,
  completed,
}) => {
  return (
    <div className="group">
      <article
        className="flex justify-between p-4 shadow-md group-hover:bg-blue-500 group-hover:text-white"
        onMouseOver={() => {
          console.log("Triggered");
        }}
      >
        <section>
          <h3 className="font-bold">{title}</h3>
          <p className="text-sm">{description}</p>
        </section>
        <input type="checkbox" name="completed" defaultChecked={completed} />
      </article>
    </div>
  );
};
