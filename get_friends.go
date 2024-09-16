package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func getFriends(driver neo4j.DriverWithContext, user string) error {
    session := driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(context.Background())

    _, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MATCH (u:User {name: $user})-[:FRIENDS_WITH]->(friend)
            RETURN friend.name
        `
        params := map[string]interface{}{
            "user": user,
        }
        result, err := tx.Run(context.Background(), query, params)
        if err != nil {
            return nil, err
        }
        for result.Next(context.Background()) {
            fmt.Println(user, "is friends with:", result.Record().Values[0])
        }
        return nil, result.Err()
    })
    return err
}