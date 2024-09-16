package main

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Recommend friends to a user (friends of friends)
func recommendFriends(driver neo4j.DriverWithContext, username string) ([]string, error) {
    session := driver.NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(context.Background())

    recommendations := []string{}

    _, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MATCH (u:User {name: $username})-[:FRIENDS_WITH]-(friend)-[:FRIENDS_WITH]-(recommendedFriend)
            WHERE NOT (u)-[:FRIENDS_WITH]-(recommendedFriend)
            RETURN recommendedFriend.name
        `
        params := map[string]interface{}{
            "username": username,
        }
        result, err := tx.Run(context.Background(), query, params)
        if err != nil {
            return nil, err
        }

        for result.Next(context.Background()) {
            recommendation := result.Record().Values[0].(string)
            recommendations = append(recommendations, recommendation)
        }
        return recommendations, result.Err()
    })

    if err != nil {
        return nil, err
    }
    return recommendations, nil
}
