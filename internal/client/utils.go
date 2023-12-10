package main

import (
	"fmt"
	"regexp"
)

func ExtractVideoId(url string) (string, error) {
	// Define a regular expression pattern for extracting YouTube video IDs
	pattern := `(?:youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]{11})`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the matches
	matches := re.FindStringSubmatch(url)

	// Check if a match is found
	if len(matches) < 2 {
		return "", fmt.Errorf("No YouTube video ID found in the URL")
	}

	// Return the video ID (captured group)
	return matches[1], nil
}
