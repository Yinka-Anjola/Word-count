package main

import (
        "fmt"
        "strings"
)

func countWord(word string) map[string]int {
        counts := map[string]int{}
        for _, text := range strings.Fields(word) {
                counts[text]++
        }
        return counts
}

func main() {
        var word string

        fmt.Println("Enter a word: ")
        fmt.Scanln(&word)

        countWords := countWord(word)

        fmt.Println("Word counts: ")
        for text, count := range countWords {
                fmt.Printf("%s: %d\n", text, count)
        }
}




