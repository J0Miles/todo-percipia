# todo-percipia
A todo application to show, create, and delete todo's

# Getting Started
**Must have mariadb or mysql installed and running daemon or instance**
Create '.ENV' file under *backend* root directory 

>Example .Env file:

>>PORT=

>>DB_HOST=

>>DB_USER=

>>DB_PASS=

## Add Task
- input text in textfield under **Item**
- Submit

## Delete Task
**THERE IS A KNOWN BUG WITH THE DELETE BUTTON - To work around this bug please see workaround section**
- click Delete button under **Action** column

## Update Task
- Input desired new name in input field under **Action** column
- Click *Update Title* button

### Workaround
- Install Postman
- add new DELETE request
- set URL to http://localhost:8000/delete/{id} *Where id is the id of the task you want to delete*
