package main

import (
	"testing"
)

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

	type testScenario struct {
		input   string
		outcome []string
	}

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
