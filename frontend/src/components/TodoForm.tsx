///mport React from "react";
///mport {Todo} from "../models/Todo";


///ype Props = {
/// todo: {
///   title: '',
///   //description: '',
///   //created_at: '',
///   //completed: '',
/// };
///;

///xport const TodoForm: React.FC<Props> = ({todo}) => {
/// return (

///   <form className="w-full max-w-sm">
///     <div className="mb-6 md:flex md:items-center">
///       <div className="md:w-1/3">
///         <label className="block pr-4 mb-1 font-bold text-gray-500 md:text-right md:mb-0">
///           {todo.title}
///         </label>
///       </div>
///       <div className="md:w-2/3">
///         <input className="w-full px-4 py-2 leading-tight text-gray-700 bg-gray-200 border-2 border-gray-200 rounded appearance-none focus:outline-none focus:bg-white focus:border-purple-500" id="inline-full-name" type="text" value="Jane Doe">
///         </input>
///       </div>
///     </div>
///     <div className="mb-6 md:flex md:items-center">
///       <div className="md:w-1/3">
///         <label className="block pr-4 mb-1 font-bold text-gray-500 md:text-right md:mb-0">
///           Password
///         </label>
///       </div>
///       <div className="md:w-2/3">
///         <input className="w-full px-4 py-2 leading-tight text-gray-700 bg-gray-200 border-2 border-gray-200 rounded appearance-none focus:outline-none focus:bg-white focus:border-purple-500" id="inline-password" type="password" placeholder="******************">
///         </input>
///       </div>
///     </div>
///     <div className="mb-6 md:flex md:items-center">
///       <div className="md:w-1/3"></div>
///       <label className="block font-bold text-gray-500 md:w-2/3">
///         <input className="mr-2 leading-tight" type="checkbox">
///           <span className="text-sm">
///             Send me your newsletter!
///           </span>
///         </input>
///       </label>
///     </div>
///     <div className="md:flex md:items-center">
///       <div className="md:w-1/3"></div>
///       <div className="md:w-2/3">
///         <button className="px-4 py-2 font-bold text-white bg-purple-500 rounded shadow hover:bg-purple-400 focus:shadow-outline focus:outline-none" type="button">
///           Sign Up
///         </button>
///       </div>
///     </div>
///   </form>
/// );

///;
import {AddTodoProps} from "../models/Todo"
import {ReactComponent as PlusIcon} from "../assets/svg/plus.svg"

export const TodoForm = ({
  handleSubmit,
  task,
  handleChange,
}: AddTodoProps) => (
  <form className="flex justify-between w-full p-4" onSubmit={handleSubmit}>
    <input
      type="text"
      name="task"
      value={task}
      placeholder={"Enter task"}
      className="flex-1 p-2 mr-2 rounded shadow text-grey-dark"
      onChange={handleChange}
    />
    <button className="w-12 h-12 p-3 rounded shadow-md hover:bg-blue-500" type="submit" aria-label="Add todo">
      <PlusIcon />
    </button>
  </form>
)
