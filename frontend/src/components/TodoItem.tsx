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
        className="flex justify-between p-6 justify-items-center shadow-md bg-white-900 group-hover:bg-blue-500 group-hover:text-white"
        onClick={() => {
          console.log("Toggle Completed");
        }}
      >
        <section>
          <h3 className="font-bold">{title}</h3>
          <p className="text-sm">{description}</p>
        </section>
        <input
          className="self-center justify-center"
          type="checkbox"
          name="completed"
          defaultChecked={completed}
        />
      </article>
    </div>
  );
};
