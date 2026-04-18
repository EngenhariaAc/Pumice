package main

import "fmt"

func main() {
    gasBubbles := 100
    lavaCooling := true

    fmt.Println("🌋 Simulating pumice formation...")

    if lavaCooling {
        gasBubbles -= 60
    }

    fmt.Printf("Remaining gas bubbles: %d\n", gasBubbles)
}
