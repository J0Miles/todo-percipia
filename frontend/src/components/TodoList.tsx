import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
import {Todo, AddTodoProps} from "../models/Todo";
import {TodoItem} from "./TodoItem";
import {TodoForm} from "./TodoForm";
import {v4 as uuidv4} from "uuid";

export const TodoList = () => {
  let data;
  console.log(new Date);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [task, setTask] = useState("");

  useEffect(() => {
    fetch('http://localhost:8001/todos', {
      method: 'GET'
    })
      .then(async res => {
        console.log(await res.body)
        data = res.json()
        console.log(data);
        setTodos(await data)
      })
    console.log(todos);
  })

  const handleAddTodo = (todo: Todo) => {
    const updatedTodos = [...todos, todo]
    setTodos(updatedTodos)
    setTask("")
  }

  const handleChange = (e: ChangeEvent) => {
    const {value} = e.target as HTMLInputElement
    setTask(value)
  }

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()

    const todo = {
      id: uuidv4(),
      title: task,
      created_at: new Date(),
      completed: false,
    }
    task && handleAddTodo(todo)
  }

  console.log(todos);

  return (
    <div className="p-4">
      <TodoForm
        handleChange={handleChange}
        handleSubmit={handleSubmit}
        task={task}
      />
      <span className="flex justify-between p-6">
        <h2 className="font-medium align-left">Task</h2>
        <h2 className="font-medium align-right">Complete</h2>
      </span>
      <ul>
        {todos.map((todo, i) => (
          <li key={i} className="p-4">
            <TodoItem
              id={todo.id}
              title={todo.title}
              description={todo?.description}
              created_at={todo.created_at}
              updated_at={todo?.updated_at}
              completed={todo.completed}
            />
          </li>
        ))}
      </ul>
    </div>
  );
};

