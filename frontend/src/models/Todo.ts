import {ChangeEvent, FormEvent} from "react"
export type Todo = {
  id: string;
  title: string;
  description?: string;
  created_at: Date;
  updated_at?: string;
  completed: boolean;
}

export type TodoProps = {
  todo: Todo
  handleCheckTodo: (id: string) => void
  handleDeleteTodo: (id: string) => void
}

export type AddTodoProps = {
  task: string
  handleSubmit: (e: FormEvent) => void
  handleChange: (e: ChangeEvent) => void
}
