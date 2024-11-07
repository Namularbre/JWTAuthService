# JWT Auth Service

This API manage users for you.

## Functionalities

- Register users
- Login user with JWT creation
- JWT validation
- The users passwords are hashed with argon2

## Setup

After cloning this repository, run this docker command to build the API image
````
docker build . -t auth-service
````
And then this command to run a container based on the image
````
docker run -d -p 8080:8080 --name auth-service-container auth-service
````

## Authors

[Namulabre](https://github.com/Namularbre)
