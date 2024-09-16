package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func createUser(driver neo4j.DriverWithContext, name string, age int) error {
    session := driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
    defer session.Close(context.Background())

    _, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            CREATE (u:User {name: $name, age: $age})
            RETURN u
        `
        params := map[string]interface{}{
            "name": name,
            "age":  age,
        }
        result, err := tx.Run(context.Background(), query, params)
        if err != nil {
            return nil, err
        }
        if result.Next(context.Background()) {
            fmt.Println("User created:", result.Record().Values[0])
        }
        return nil, result.Err()
    })
    return err
}