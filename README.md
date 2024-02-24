## REST API Endpoints

### Access Home Route
- **URI:** `http://localhost:8080/home`
- **Method:** GET

### Get All Users
- **URI:** `http://localhost:8080/users`
- **Method:** GET

### Create a User
- **URI:** `http://localhost:8080/users`
- **Method:** POST
- **Body:** `{ "name": "Subhradep Ray" }`

### Get a User by ID
- **URI:** `http://localhost:8080/users/{id}`
- **Method:** GET
- **Note:** Replace `{id}` with the actual ID of the user.

### Update a User by ID
- **URI:** `http://localhost:8080/users/{id}`
- **Method:** PUT
- **Body:** `{ "name": "Tufan" }`
- **Note:** Replace `{id}` with the actual ID of the user.

### Delete a User by ID
- **URI:** `http://localhost:8080/users/{id}`
- **Method:** DELETE
- **Note:** Replace `{id}` with the actual ID of the user.
