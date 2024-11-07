# JWT Auth Service

This API manage users for you.

## Functionalities

- Register users
- Login user with JWT creation
- JWT validation
- The users passwords are hashed with argon2

## Setup

After cloning this repository, create a file called **auth.db** and another called **.env**

Fill the .env with this template, and configure the application as you like
````dotenv
# Set the database drive name used by go
DB_DRIVER_NAME=sqlite3
# Set the environment variable for the MariaDB connection string
CONN_STR="./auth.db"
# Set the environment variable for the http port
PORT=8080
# Set the environment variable for the http address
ADDRESS=0.0.0.0
# Set the salt for password hasing
SALT=0çAzà)'(-
# Set the secret key for jwt tokens
SECRET_KEY=secret_key_you_should_not_give
# Set the gin web framework env (debug release or test)
GIN_MODE=debug
````

Run this docker command to build the API image
````bash
docker build . -t auth-service
````
And then this command to run a container based on the image
````bash
docker run -d -p 8080:8080 --name auth-service-container auth-service
````

## Authors

[Namulabre](https://github.com/Namularbre)
