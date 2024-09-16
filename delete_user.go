package main

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Delete a user and their relationships
func deleteUser(driver neo4j.DriverWithContext, username string) error {
    session := driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
    defer session.Close(context.Background())

    _, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MATCH (u:User {name: $username})
            DETACH DELETE u
        `
        params := map[string]interface{}{
            "username": username,
        }
        _, err := tx.Run(context.Background(), query, params)
        return nil, err
    })

    return err
}
