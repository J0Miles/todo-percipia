import React, {ChangeEvent, FormEvent, useState, useEffect} from "react";
import {Todo, AddTodoProps} from "../models/Todo";
import {TodoItem} from "./TodoItem";
import {TodoForm} from "./TodoForm";

export const TodoList = () => {
  let data: any = {}
  console.log(new Date);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [task, setTask] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [isComplete, setCompleted] = useState<boolean>(false)
  let lastId = todos && todos[todos.length - 1];
  console.log("!!!!!!!LASTID!!!!!: ", lastId);
  console.log(task, description)
  console.log("Todo length", todos?.length)
  console.log(isComplete);
  console.log(todos)

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
  }, [])

  const handleAddTodo = (todo: any) => {
    const updatedTodos = [...todos, todo]
    setTodos(updatedTodos)
    setTask("")
    setDescription("")
  }

  const handleSubmit = (e: FormEvent) => {
    if (task !== "") {
      const todo = {
        title: task,
        created_at: new Date(),
        Completed: false,
      }

      fetch('/add', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({
          "title": task,
          "description": description,
          "completed": false
        })
      })
      task && handleAddTodo(todo)

    }
  }


  const handleDescription = (e: ChangeEvent) => {
    const {value} = e.target as HTMLInputElement
    setDescription(value)
  }

  const deleteTask = (id: number) => {
    console.log(id)
    fetch(`/delete/${id}`)
      .then(() => window.location.reload())
  }

  const handleChange = (e: ChangeEvent) => {
    const {value} = e.target as HTMLInputElement
    setTask(value)
    setCompleted(!isComplete)
  }

  const toggleComplete = (id: any) => {
    console.log(id)
    setCompleted(!isComplete)
    fetch(`/complete/${id}`)
      .then(() => window.location.reload())
  }

  return (
    <div className="p-4">
      <TodoForm
        handleDescription={handleDescription}
        handleChange={handleChange}
        handleSubmit={handleSubmit}
        task={task}
        description={description}
      />
      <span className="flex justify-between p-6">
        <h2 className="font-medium align-left">Task</h2>
        <h2 className="font-medium align-right">Complete</h2>
      </span>
      <ul>
        {todos && todos.map((todo, i) => (
          < li key={i} className="p-4" >
            <TodoItem
              Id={todo?.Id}
              title={todo.title}
              description={todo?.description}
              created_at={todo.created_at}
              updated_at={todo?.updated_at}
              Completed={todo.Completed}
              toggleComplete={toggleComplete}
              deleteTask={deleteTask}
            />
          </li>
        ))}
      </ul>
    </div >
  );
};

