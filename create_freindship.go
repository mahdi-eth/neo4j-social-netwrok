package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func createFriendship(driver neo4j.DriverWithContext, user1 string, user2 string) error {
    session := driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
    defer session.Close(context.Background())

    _, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MATCH (u1:User {name: $user1}), (u2:User {name: $user2})
            CREATE (u1)-[:FRIENDS_WITH]->(u2)
            RETURN u1, u2
        `
        params := map[string]interface{}{
            "user1": user1,
            "user2": user2,
        }
        result, err := tx.Run(context.Background(), query, params)
        if err != nil {
            return nil, err
        }
        if result.Next(context.Background()) {
            fmt.Println("Friendship created between", user1, "and", user2)
        }
        return nil, result.Err()
    })
    return err
}
