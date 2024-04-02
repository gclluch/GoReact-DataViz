# GoReact-DataViz
Basic data visualizer using Go and React

### Running the Application with Docker Compose

1. **Clone the Repository**

First, clone this repository to your local machine:
```sh
git clone https://github.com/gclluch/GoReact-DataViz
cd GoReact-DataViz
```

2. **Start the Application**

Use Docker Compose to build and start the application:

```sh
docker-compose up --build
```

After the containers are up and running, the frontend will be accessible at http://localhost:3000, and the backend will be available at http://localhost:8080.


### Running tests

1. **Backend**

Navigate to backend directory and run tests with coverage:
```sh
cd backend
go test -coverprofile=coverage.out
```

2. **Frontend**
Navigate to frontend directory, install dependencies and run tests:

```sh
cd frontend
npm install
npm test -- --watchAll=false
```


