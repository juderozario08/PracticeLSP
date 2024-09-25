package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"practiceLSP/lsp"
	"practiceLSP/rpc"
)

func main() {
	logger := getLogger("/home/juderozario/PracticeLSP/log.txt")
	logger.Println("LSP Attached")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.Decoder(msg)
		if err != nil {
			logger.Printf("Got an error %v", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received message with method %v", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Not able to parse this")
		}
		logger.Printf("Connected to %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)
		reply := rpc.Encoder(lsp.NewInitializeResponse(request.ID))
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Print("Sent reply")
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o744)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[practiceLSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
