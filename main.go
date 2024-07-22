package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// git hub repo details
	owner := "parveshbhardwaj"
	repo := "workflow-dispatcher"
	workflowID := "main.yaml"

	// creating work flow input
	workflowInput := struct {
		Ref string `json:"ref"`
	}{Ref: "main"}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(workflowInput)
	if err != nil {
		log.Fatal(err)
		return
	}

	// creating workflow dispatcher request
	workflowDispatchURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/workflows/%s/dispatches",
		owner, repo, workflowID)

	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, workflowDispatchURL, &buf)

	// setting header to the request
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer <Access Token>")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	if err != nil {
		log.Printf("error building the request: %v", err)
		return
	}
	// send the workflow dispatcher request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("workflow triggered with status code %s \n", resp.Status)
}
