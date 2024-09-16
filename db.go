package main

import (
    "log"

    "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func connectToNeo4j() neo4j.DriverWithContext {
    uri := "neo4j://neo4j:7687"
    username := "neo4j"
    password := "test1234"

    driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
    if err != nil {
        log.Fatal("Error connecting to Neo4j:", err)
    }
    return driver
}