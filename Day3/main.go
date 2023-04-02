package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("input3.txt")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        // split each line in two equal parts of string
        // and find common characters and save them in a slice
        var common []string

        //convert findCommonChar function output to string and append it to common slice
        common = append(common, string(findCommonChar(line[:len(line)/2], line[len(line)/2:])))

        // replace a-z with 1-26 and A-Z with 27-52 in common slice and add them in sum
        for _, v := range common {
            if v >= "a" && v <= "z" {
                sum += int(v[0]) - 96
            } else if v >= "A" && v <= "Z" {
                sum += int(v[0]) - 38
            }
        }

    }
    fmt.Println(sum)
}

func findCommonChar(str1 string, str2 string) rune {
    for _, ch := range str1 {
        if strings.ContainsRune(str2, ch) {
            return ch
        }
    }
    return 0 // If no common character is found
}
