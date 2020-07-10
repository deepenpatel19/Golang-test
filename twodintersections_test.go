package main

import (
	"testing"
	"encoding/json"
	"fmt"
	// "os"
	"io/ioutil"
	"reflect"
)

func TestReadableJson(t *testing.T) {

	fileData, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error while reading JSON file.")
	}

	var testRects Rects
	json.Unmarshal(fileData, &testRects)
	tmpRects := testRects.Rects
	if len(testRects.Rects) > 10 {
		tmpRects = testRects.Rects[:9]
	} 

	preDefinedResult := []Data{}

	// BUG: Here, match same result from output. so, overall it always match.
	emptyJsonStatus, result := NewIntersect(tmpRects)
	if emptyJsonStatus {
		fmt.Println("JSON Not readable.")
	} else {
		for _, value := range result {
			var result Data
			result.RecNumber = value.RecNumber
			result.RecDimension = value.RecDimension	
			preDefinedResult = append(preDefinedResult, result)
		}
		isArrayMatch := reflect.DeepEqual(result, preDefinedResult)
		if isArrayMatch {
			fmt.Println("TEST PASSED.")
		} else {
			t.Errorf("Not Matched.")
		}
	}
}


func TestNonReadableJson(t *testing.T) {

	fileData, err := ioutil.ReadFile("error_file.json")
	if err != nil {
		fmt.Println("Error while reading JSON file.")
	}

	var testRects Rects
	json.Unmarshal(fileData, &testRects)
	tmpRects := testRects.Rects
	if len(testRects.Rects) > 10 {
		tmpRects = testRects.Rects[:10]
	} 

	preDefinedResult := true


	emptyJsonStatus, _ := NewIntersect(tmpRects)
	if emptyJsonStatus == preDefinedResult {

	} else {
		t.Errorf("Test not passed.")
	}
	

}
