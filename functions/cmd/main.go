package main

import (
	"fmt"
	"log"
	"os"

	functions "upload-token.functions"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	funcframework.RegisterHTTPFunction("/challenge", functions.ChallengeHandler)
	funcframework.RegisterHTTPFunction("/submitReport", functions.SubmitReportHandler)

	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	fmt.Println("Listening port:", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
