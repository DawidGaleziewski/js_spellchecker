package main

import (
	"testing"
)

type testScenario struct {
	input   string
	outcome []string
}

const testBlob1 = `import React from 'react';
import styled from '@emotion/styled';
import {css} from '@emotion/react';

interface IBackground {
    variant: "light" | "dark"
}

export const Background = styled.div<IBackground>

export const testVariable
function testingNameOfVariable(){`

func TestFindVariableNames(t *testing.T) {

	var testScenarios = []testScenario{
		testScenario{testBlob1, []string{"testVariable", "Background"}},
	}

	for _, testSuite := range testScenarios {
		result := GetVariableNames(testSuite.input)
		if !containsEvery(result, testSuite.outcome) {
			t.Error("expected for testing input: ", testSuite.input, ",to find all declarations: ", testSuite.outcome)
		}
	}
}

func TestCamelCaseToWords(t *testing.T) {
	var testScenarios = []testScenario{
		testScenario{"test", []string{"test"}},
		testScenario{"Test", []string{"test"}},
		testScenario{"testVariable", []string{"test", "variable"}},
		testScenario{"TestVariable", []string{"test", "variable"}},
		testScenario{"TestVariableSuffix", []string{"test", "variable", "suffix"}},
	}

	for _, testSuite := range testScenarios {
		result := CamelCaseToWords(testSuite.input)
		if !containsEvery(result, testSuite.outcome) {
			t.Error("expected for testing input: ", testSuite.input, ",to find all words: ", testSuite.outcome, "Found only: ", result)
		}
	}
}

func containsOne(arr []string, searchItem string) bool {
	for _, item := range arr {
		if item == searchItem {
			return true
		}
	}
	return false
}

func containsEvery(arr []string, searchItems []string) bool {
	var foundItems []string
	numberOfSearchedItems := len(searchItems)

	for _, searchItem := range searchItems {
		if containsOne(arr, searchItem) {
			foundItems = append(foundItems, searchItem)
		}
	}

	return numberOfSearchedItems == len(foundItems)
}
