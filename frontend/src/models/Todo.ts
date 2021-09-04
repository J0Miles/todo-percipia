import {ChangeEvent, FormEvent} from "react"
export type Todo = {
  Id: number;
  title: string;
  description?: string;
  created_at: Date;
  updated_at?: string;
  Completed: boolean;
  toggleComplete?: any;
  deleteTask?: (id: number) => void;
}

export type TodoProps = {
  todo: Todo
  handleCheckTodo: (id: string) => void
  handleDeleteTodo: (id: string) => void
}

export type AddTodoProps = {
  task: string
  description: string
  handleSubmit: (e: FormEvent) => void
  handleChange: (e: ChangeEvent) => void
  handleDescription: (e: ChangeEvent) => void
}

export type UpdateTodoProps = {
  id: number
  setCompleted: () => void;
}
