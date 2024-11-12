package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// Ensure proper number of arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <variables_file> <template_file> <output_file>")
		return
	}

	// Get the file paths from command-line arguments
	variablesFile := os.Args[1]
	templateFile := os.Args[2]
	outputFile := os.Args[3]

	// Read the variables from the provided variables file
	vars, err := readVariables(variablesFile)
	if err != nil {
		fmt.Printf("Error reading variables from '%s': %v\n", variablesFile, err)
		return
	}

	// Read the Markdown template file
	tmplContent, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Printf("Error reading template file '%s': %v\n", templateFile, err)
		return
	}

	// Create the template and execute with the variables map
	tmpl, err := template.New("markdown").Parse(string(tmplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Apply the template to the variables map and write the result to the specified output file
	var output bytes.Buffer
	err = tmpl.Execute(&output, vars)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Save the output to the specified output file
	err = os.WriteFile(outputFile, output.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file '%s': %v\n", outputFile, err)
		return
	}
}

// readVariables reads a text file containing key=value pairs and returns a map
func readVariables(filePath string) (map[string]string, error) {
	vars := make(map[string]string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore empty lines or lines that don't contain '='
		if strings.TrimSpace(line) == "" || !strings.Contains(line, "=") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			vars[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return vars, nil
}
