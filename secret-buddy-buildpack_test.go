package main

import (
	"os"
	"testing"
)

func TestExportEnvVarsFromMap(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"DB_PASSWORD": "current.DB_PASSWORD,previous.DB_PASSWORD",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current_password,previous_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestExportEnvVarsFromMapEmptyRules(t *testing.T) {
	// Define a test case
	testCase := struct {
		inputEnv       string
		inputRules     map[string]string
		expectedOutput map[string]string
	}{
		inputEnv: `{"current": {"DB_PASSWORD": "current_password"}, "previous": {"DB_PASSWORD": "previous_password"}}`,
		inputRules: map[string]string{
			"": "",
		},
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current_password",
		},
	}

	// Call the ExportEnvVarsFromMap function with the test case input
	outputValue, err := ExportEnvVarsFromMap(testCase.inputEnv, testCase.inputRules)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestParseRules(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput map[string]string
	}{
		input: `{"DB_PASSWORD": "current.DB_PASSWORD"}`,
		expectedOutput: map[string]string{
			"DB_PASSWORD": "current.DB_PASSWORD",
		},
	}

	// Call the ParseRules function with the test case input
	outputValue, err := ParseRules(testCase.input)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestParseEmptyRules(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput map[string]string
	}{
		input: "{}",
		expectedOutput: map[string]string{
			"": "",
		},
	}

	// Call the ParseRules function with the test case input
	outputValue, err := ParseRules(testCase.input)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	for key, value := range testCase.expectedOutput {
		if outputValue[key] != value {
			t.Errorf("Expected %s but got %s for key %s", value, outputValue[key], key)
		}
	}
}

func TestGetEnvVar(t *testing.T) {
	// Define a test case
	testCase := struct {
		input          string
		expectedOutput string
	}{
		input:          "SECRETBUDDY_ENV",
		expectedOutput: "test_value",
	}

	// Set the environment variable for the test case
	err := os.Setenv(testCase.input, testCase.expectedOutput)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Call the GetEnvVar function with the test case input
	outputValue, err := GetEnvVar(testCase.input)

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the output matches the expected output
	if outputValue != testCase.expectedOutput {
		t.Errorf("Expected %s but got %s", testCase.expectedOutput, outputValue)
	}

}
