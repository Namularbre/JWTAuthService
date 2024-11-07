FROM golang:1.23

# we work in an app directory
WORKDIR /app

# we move go.mod & go.sum in the app dir
COPY go.* ./

# we download the dependancies
RUN go mod download
# we create a vendor directory to store the dependencies localy
RUN go mod vendor
# we verify the dependencies
RUN go mod verify

# we copy the rest of the files
COPY . .

# we create an executable
RUN go build -o authService

# we expose the port used by the API
EXPOSE 8080

# we run the executable
CMD ["./authService"]
