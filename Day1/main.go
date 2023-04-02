package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func main() {
    // Open the file containing the Elves' inventory
    file, err := os.Open("Input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Create a scanner to read the file line-by-line
    scanner := bufio.NewScanner(file)

    // Create a map to store the total Calories for each Elf
    elfCalories := make(map[int]int)

    // Initialize some variables to keep track of the current Elf and their Calories
    currentElf := 0
    currentCalories := 0

    // Loop through each line in the file
    for scanner.Scan() {
        // Read the current line into a variable
        line := scanner.Text()

        // If the line is blank, move on to the next Elf
        if line == "" {
            elfCalories[currentElf] += currentCalories
            currentElf++
            currentCalories = 0
        } else {
            // Otherwise, try to convert the line to an integer and add it to the current Elf's Calories
            if num, err := strconv.Atoi(line); err == nil {
                currentCalories += num
            }
        }
    }

    // Add the final Elf's Calories to the map
    elfCalories[currentElf] += currentCalories

    // Create a slice of the Elves' total Calories
    calorieSlice := make([]int, 0, len(elfCalories))
    for _, calories := range elfCalories {
        calorieSlice = append(calorieSlice, calories)
    }

    // Sort the slice of Calories in descending order
    sort.Sort(sort.Reverse(sort.IntSlice(calorieSlice)))

    // Print the total Calories for the top three Elves
    var totalCalories int
    for i, calories := range calorieSlice {
        if i < 3 {
            totalCalories += calories
            fmt.Printf("Elf %d has %d Calories\n", i+1, calories)
        } else {
            break
        }
    }
    fmt.Printf("The top three Elves are carrying a total of %d Calories\n", totalCalories)

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }
}
