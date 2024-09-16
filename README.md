# Neo4j-social-network

## 🧑‍🤝‍🧑 A Social Network Built with Go and Neo4j Graph Database

This project demonstrates how to build a simple **social network** using **Go** as the backend language and **Neo4j** as the graph database to handle relationships between users. The application includes features like mutual friend discovery, friend recommendations, and user management. The whole project is containerized using Docker, making it easy to deploy and run.

---

## 🎯 Features

- **Mutual Friends**: Find mutual friends between two users.
- **Friend Recommendations**: Suggest friends based on friends of friends.
- **User Management**: Add, delete, and query users.
- **Neo4j Integration**: Efficient graph data modeling using Neo4j.
- **Dockerized**: Run with Docker and Docker Compose for easy setup and deployment.

---

## 🚀 Getting Started

Follow the steps below to get the project up and running.

### Prerequisites

Make sure you have the following installed:

- **Go** (1.18 or higher)
- **Docker** and **Docker Compose**

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mahdi-eth/neo4j-social-network.git
   cd neo4j-social-network
   ```

2. **Set up Docker Compose**: Ensure Neo4j and the Go application run seamlessly using Docker Compose.
   
 ```bash
  docker-compose up --build
  ```

3. **Access the Neo4j Browser**: Open your browser and go to http://localhost:7474. Use these credentials:

- **Username**: neo4j
- **Password**: test1234


## 🛠️ Usage

### See the network graph:

```go
  MATCH (u:User) RETURN u;
```


