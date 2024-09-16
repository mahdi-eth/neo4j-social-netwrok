package main

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Find mutual friends between two users
func findMutualFriends(driver neo4j.DriverWithContext, user1, user2 string) ([]string, error) {
    session := driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(context.Background())

    friends := []string{}

    _, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MATCH (u1:User {name: $user1})-[:FRIENDS_WITH]-(mutualFriend)-[:FRIENDS_WITH]-(u2:User {name: $user2})
            RETURN mutualFriend.name
        `
        params := map[string]interface{}{
            "user1": user1,
            "user2": user2,
        }
        result, err := tx.Run(context.Background(), query, params)
        if err != nil {
            return nil, err
        }

        for result.Next(context.Background()) {
            friend := result.Record().Values[0].(string)
            friends = append(friends, friend)
        }
        return friends, result.Err()
    })

    if err != nil {
        return nil, err
    }
    return friends, nil
}
