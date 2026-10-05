// Word Frequency Counter

package main

import (
    "fmt"
    "os"
    "path/filepath"
    "bufio"
    "slices"
    "os/signal"
    "syscall"
)

func main() {
    // Use traditional unix pipe behavior
    signal.Ignore(syscall.SIGPIPE)

    if len(os.Args) < 2 {
        fmt.Println("Usage:", filepath.Base(os.Args[0]), "<filename>")
        return 
    }

    // read file line by line
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    wordMap := make(map[string]int)

    // Create scanner to scan words
    scanner := bufio.NewScanner(file)

    // Tell scanner to split by words
    scanner.Split(bufio.ScanWords)

    // Loop over word in text
    for scanner.Scan() {
//        fmt.Println(scanner.Text())
        word := scanner.Text()
        wordMap[word]++
//        fmt.Println(word)
    }

    // Check for scan errors
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }

/*
    testMap := make(map[string]int)

    testMap["fart"] = 7
    testMap["blah"] = 1
    testMap["darn"] = 5
    testMap["shut"] = 8
    testMap["boom"] = 77
    testMap["doop"] = 7

    counts, words := sortWordCount(testMap)
    fmt.Println(counts)
    fmt.Println(words)
*/
    counts, words := sortWordCount(wordMap)

    for i := 0; i < len(words); i++ {
        fmt.Printf("Word: %s\nCount: %d\n\n", words[i], counts[i])
    }
}

func sortWordCount(wordMap map[string]int) ([]int, []string) {
    var words []string
    var counts []int

    for word, count := range wordMap {
        // append value to count slice
        counts = append(counts, count)
        // append value to words slice
        words = append(words, word)

    }

//    fmt.Println(words)
//    fmt.Println(counts)

    total := len(counts)

    // Bubble Sort
    for {
        sorts := 0
        for i, value := range counts {
//            fmt.Println(value)

            // Done cooking
            if i == total-1 { break }

            // Swap them
            n := i+1
            if value > counts[n] {
                sorts++
                counts[i], counts[n] = counts[n], counts[i]
                words[i], words[n] = words[n], words[i]
            }
        }
        if sorts == 0 { break }
    }

    slices.Reverse(counts)
    slices.Reverse(words)

    return counts, words
}


