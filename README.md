# go-todo-rest-api
To-Do list REST API
URL: https://todo-app-exampl.herokuapp.com

### Endpoints:
for /api methods JWT token required

endpoint | method | body | response
--- | --- | --- | ---
/auth/sign-up | POST | { name, username, password} | user's id
/auth/sign-in | POST | { username, password} | JWT token
/api/lists | GET | - | all user's lists
/api/lists | POST | { title, description (optional) } | list's id
/api/lists/:id | GET | - | list by id
/api/lists/:id | PUT | { title, description (optional) } | status of updating
/api/lists/:id | DELETE | | status of deletion
/api/lists/:id/items | GET | - | items of list by id
/api/lists/:id/items | POST | { title, description (optional), done (boolean, optional) } | items of list by id
/api/items/:id | GET | - | item by id
/api/items/:id | PUT | { title, description (optional), done (boolean, optional) } | status of updating
/api/items/:id | DELETE | - | status of deletion
