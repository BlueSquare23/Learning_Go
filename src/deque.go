package main

import (
	"encoding/json"
	"errors"
	"flag" // TODO: Replace this with pflag https://pkg.go.dev/github.com/spf13/pflag
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Opt struct {
	QueueFile string
	Backend   string
}

type ArgV struct {
	Queue     string
	Operation string
	Data      string
}

func main() {
	opt := parseFlags()
	args := parseArgV()

	var queueMap map[string]any

	jsonBytes := readQueueFile(opt.QueueFile)

	//    fmt.Printf("Json File: %s, Bytes: %s\n", opt.QueueFile, jsonBytes)

	if len(jsonBytes) != 0 {
		if err := json.Unmarshal(jsonBytes, &queueMap); err != nil {
			log.Fatalf("Failed to unpack json: %s", err)
		}
	}

	dq, ok := queueMap[args.Queue].([]any)
	if !ok {
		// Error if unshift or pop
		if args.Operation == "shift" || args.Operation == "pop" {
			log.Fatalf("Queue " + args.Queue + " not found")
		}

		// Create it for unshift or push
		queueMap[args.Queue] = []any{}
	}
	dq = queueMap[args.Queue].([]any)

	switch args.Operation {
	case "list":
		fmt.Println(strings.Join(dq, ", "))
	case "shift":
		fmt.Println(shift(&dq))
	case "unshift":
		unshift(&dq, args.Data)
	case "pop":
		fmt.Println(pop(&dq))
	case "push":
		push(&dq, args.Data)
	default:
		fmt.Fprintln(os.Stderr, "Invalid Operation:", "'"+string(args.Operation)+"'",
			"Valid Operators: shift, unshift, pop, push")
		os.Exit(1)
	}

	queueMap[args.Queue] = dq
	writeQueueFile(toJson(queueMap), opt.QueueFile)
}

func parseFlags() Opt {
	var opt Opt

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	queueFile := filepath.Join(homeDir, ".dq.json")

	flag.StringVar(&opt.QueueFile, "file", queueFile, "File to store dequeues in")
	// TODO: Implement other storage backends
	//    flag.StringVar(&opt.Backend, "backend", 'file', "Store backend (file, redis, db)")
	flag.Parse()
	return opt
}

func parseArgV() ArgV {
	// Rest of args after flag.Args parsed.
	argv := flag.Args()

	scriptName := filepath.Base(os.Args[0])
	if len(argv) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:", scriptName, "<queue> <operation> [<data>]")
		os.Exit(1)
	}

	args := ArgV{
		argv[0],
		argv[1],
		"",
	}

	if len(argv) > 2 {
		args.Data = argv[2]
	}

	return args
}

func readQueueFile(queueFile string) []byte {
	content, err := os.ReadFile(queueFile)
	if err != nil {
		// Create file if it doesn't exist
		if errors.Is(err, os.ErrNotExist) {
			file, createErr := os.OpenFile(queueFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
			if createErr != nil {
				log.Fatalf("Failed to create file: %v\n", createErr)
			}

			file.Close()

			// Data is empty since we just created it
			content = []byte{}
		} else {
			log.Fatalf("Failed to read file: %s", err)
		}
	}
	return content
}

func shift(dq *[]any) string {
	first := (*dq)[0]
	*dq = (*dq)[1:]
	return first.(string)
}

func unshift(dq *[]any, data any) {
	*dq = append([]any{data}, *dq...)
	fmt.Println(dq)
}

func pop(dq *[]any) string {
	last := (*dq)[len(*dq)-1]
	*dq = (*dq)[:len(*dq)-1]
	return last.(string)
}

func push(dq *[]any, data any) {
	*dq = append(*dq, data)
}

func toJson(q map[string]any) string {
	// Convert map to JSON bytes
	jsonBytes, err := json.Marshal(q)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	// Convert bytes to string and return it
	jsonString := string(jsonBytes)
	return jsonString
}

func writeQueueFile(json, queueFile string) {
	err := os.WriteFile(queueFile, []byte(json), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
