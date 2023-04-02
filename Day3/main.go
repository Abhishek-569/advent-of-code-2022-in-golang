package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Open the input file for reading
    file, err := os.Open("input3.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Keep track of the current group of lines being processed
    var lines []string
    totalSum := 0

    // Read the lines from the file and group them into sets of 3
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
        if len(lines) == 3 {
            // Find the common character in the current group of lines
            common := findCommonCharacter(lines[0], lines[1], lines[2])
            //for a-z add 1-26 and for A-Z add  27-52 add into total sum
            for _, r := range common {
                if r >= 'a' && r <= 'z' {
                    totalSum += int(r - 'a' + 1)
                } else if r >= 'A' && r <= 'Z' {
                    totalSum += int(r - 'A' + 27)
                }
            }

            // Reset the lines buffer for the next group
            lines = []string{}
        }
    }

    // Check for any errors that occurred during scanning
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    fmt.Println(totalSum)
}

// Helper function to find the common character in 3 strings
func findCommonCharacter(s1, s2, s3 string) string {
    // Convert the strings to sets of runes (characters)
    set1 := make(map[rune]bool)
    for _, r := range s1 {
        set1[r] = true
    }
    set2 := make(map[rune]bool)
    for _, r := range s2 {
        set2[r] = true
    }
    set3 := make(map[rune]bool)
    for _, r := range s3 {
        set3[r] = true
    }

    // Find the intersection of the 3 sets
    for r := range set1 {
        if set2[r] && set3[r] {
            return string(r)
        }
    }
    // If no common character was found, return an empty string
    return ""
}
