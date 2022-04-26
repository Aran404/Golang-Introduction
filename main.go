package main

// Imports
import (
    "fmt"
    "time"
    "math/rand"
    "strings"
)

var letters = []rune("rps")

// Gets computers input
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


func main() {
    rand.Seed(time.Now().UnixNano())
    
    // Gets user input
    fmt.Print("Rock (r) paper (p) or scissors (s) > ")
    var user_input string
    fmt.Scanln(&user_input)
    lower_user_input := strings.ToLower(user_input)
    
    // Checks to see if input is not r , p or s
    if lower_user_input == "r" || lower_user_input == "p" || lower_user_input == "s"{
    } else {
        fmt.Println("Answer has to be rock paper or scissors")
        return
    }

    // Computer
    computer_input := randSeq(1)
    
    // Displays choices from both computer and user
    fmt.Printf("Computer Picked: %s, You Picked: %s\n", computer_input, lower_user_input)
    
    // Checks who wins
    if computer_input == lower_user_input {
        fmt.Println("Draw Game")
    } else if computer_input == "s" && lower_user_input == "p" {
        fmt.Println("Computer Wins")
    } else if computer_input == "s" && lower_user_input == "r" {
        fmt.Println("You Win")
    } else if computer_input == "p" && lower_user_input == "r" {
        fmt.Println("Computer Wins")
    } else if computer_input == "p" && lower_user_input == "s" {
        fmt.Println("You Win")
    } else if computer_input == "r" && lower_user_input == "s" {
        fmt.Println("Computer Wins")
    } else if computer_input == "r" && lower_user_input == "p" {
        fmt.Println("You Win")
    }
    
}
