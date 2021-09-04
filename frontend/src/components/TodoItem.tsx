import React from "react";
import {Todo} from "../models/Todo";
import {ReactComponent as DeleteButton} from "../assets/svg/icons8-delete.svg"

export const TodoItem: React.FC<Todo> = ({
  Id,
  title,
  description,
  created_at,
  updated_at,
  Completed,
  toggleComplete,
  deleteTask
}) => {
  let content;
  content = Completed ? (
    <>
      <h3 className="font-bold line-through">{title}</h3>
      <p className="text-sm line-through">{description}</p>
    </>
  ) : (
    <>
      <h3 className="font-bold">{title}</h3>
      <p className="text-sm">{description}</p>
    </>);
  return (
    <div className="group">
      <article
        className="flex justify-between p-6 shadow-md justify-items-center bg-white-900 group-hover:bg-blue-500 group-hover:text-white"
      >
        <section>
          {content}
        </section>
        <div className="flex items-center content-center self-center justify-center row-end-2">
          <input
            className="self-center justify-center w-5 h-5"
            type="checkbox"
            name="completed"
            defaultChecked={Completed}
            onClick={() => toggleComplete(Id)}
          />
          &nbsp;
          <div className="flex">
            <button className="items-center justify-center w-5 h-5">
              <DeleteButton className="w-5 h-5" onClick={() => deleteTask && deleteTask(Id)} />
            </button>
          </div>
        </div>
      </article>
    </div >
  );
}
