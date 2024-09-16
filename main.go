package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
    driver := connectToNeo4j()
    defer driver.Close(context.Background())

    createUser(driver, "Alice", 30)
    createUser(driver, "Bob", 25)
    createFriendship(driver, "Alice", "Bob")
    createUser(driver, "Charlie", 28)
    createFriendship(driver, "Bob", "Charlie")

    // Find mutual friends
    mutualFriends, err := findMutualFriends(driver, "Alice", "Bob")
    if err != nil {
        log.Fatal("Error finding mutual friends:", err)
    }
    fmt.Println("Mutual friends between Alice and Bob:", mutualFriends)

    // Recommend friends to Alice
    recommendations, err := recommendFriends(driver, "Alice")
    if err != nil {
        log.Fatal("Error recommending friends:", err)
    }
    fmt.Println("Friend recommendations for Alice:", recommendations)
    

    if err := getFriends(driver, "Alice"); err != nil {
        log.Fatal("Error fetching friends:", err)
    }
}
