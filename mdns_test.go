package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestDiscovery_Run(t *testing.T) {
	// Create a temporary file for output.
	outputFile, err := ioutil.TempFile("", "output.*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outputFile.Name())

	// Set up a Discovery with a short interval for testing.
	d := &Discovery{
		interval: 100 * time.Millisecond,
	}

	// Set up a context that will cancel after a short time.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Set up a channel to receive the TargetGroups.
	ch := make(chan []*TargetGroup)

	// Start the Discovery running.
	go d.Run(ctx, ch)

	// Wait for the first set of results to arrive.
	select {
	case tg := <-ch:
		// Check that the output is valid JSON and contains at least one target.
		if len(tg) == 0 {
			t.Error("Expected at least one target group, got none")
		}
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(tg); err != nil {
			t.Errorf("Error encoding TargetGroups as JSON: %v", err)
		}
		if _, err := json.Marshal(tg); err != nil {
			t.Errorf("Error marshaling TargetGroups as JSON: %v", err)
		}
		if !json.Valid(buf.Bytes()) {
			t.Errorf("Output is not valid JSON: %s", buf.String())
		}

		// Check that the output was written to the file.
		outputBytes, err := ioutil.ReadFile(outputFile.Name())
		if err != nil {
			t.Errorf("Error reading output file: %v", err)
		}
		if !bytes.Equal(buf.Bytes(), outputBytes) {
			t.Errorf("Output to file does not match output to console:\n%s\n%s", buf.String(), string(outputBytes))
		}
	case <-ctx.Done():
		t.Fatalf("Discovery did not produce any results before timeout: %v", ctx.Err())
	}
}

func TestTargetGroups_Less(t *testing.T) {
	tg1 := &TargetGroup{
		Targets: []string{"b.example.com:1234", "a.example.com:1234"},
	}
	tg2 := &TargetGroup{
		Targets: []string{"c.example.com:1234", "d.example.com:1234"},
	}
	tg3 := &TargetGroup{
		Targets: []string{"d.example.com:1234", "c.example.com:1234"},
	}
	tg4 := &TargetGroup{
		Targets: []string{},
	}
	targetGroups := TargetGroups{tg2, tg1, tg3, tg4}

	// Sort the TargetGroups.
	sort.Sort(&targetGroups)

	// Check that tg1, tg2, tg3, tg4 are in the correct order.
	expectedOrder := []*TargetGroup{tg1, tg2, tg3, tg4}
	if !reflect.DeepEqual(targetGroups, expectedOrder) {
		t.Errorf("TargetGroups are not in the expected order:\n%#v\n%#v", targetGroups, expectedOrder)
	}
}
