package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	// "fmt"
	//  s "strings"
)

func TestGetAllThreads(t *testing.T) {
	req, err := http.NewRequest("GET", "/Threads/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler ret urned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `[{"Id":1,"Discription":"My Discription","Title":"My title","TimeCreated":"15:19:19, Jun 10 2019"},{"Id":2,"Discription":"My Discription2","Title":"My title2","TimeCreated":"15:19:19, Jun 10 2019"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}

func TestGetSingleThread(t *testing.T) {
	req, err := http.NewRequest("GET", "/Threads/1", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler ret urned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `{"Id":1,"Discription":"My Discription","Title":"My title","TimeCreated":"15:19:19, Jun 10 2019"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestGetUnexistingSingleThread(t *testing.T) {
	req, err := http.NewRequest("GET", "/Threads/3", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler ret urned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `Id not Found`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestGetInvalidSingleThread(t *testing.T) {
	req, err := http.NewRequest("GET", "/Threads/3s", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler ret urned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `There was an error Processing your request`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}

func TestDeleteThread(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/Threads/2", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler ret urned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `Thread Deleted Successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestDeleteUnexistingThread(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/Threads/3", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler ret urned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `Id not Found`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestDeleteIvalidThread(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/Threads/3s", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
	expected := `There was an error Processing your request`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestPostThread(t *testing.T) {

	var jsonStr = []byte(`{
		"Id": 3,
		"Discription": "My aaaaaa",
		"Title": "My title"
	  }`)

	req, err := http.NewRequest("POST", "/Thread", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Thread Added Successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestExistingPostThread(t *testing.T) {

	var jsonStr = []byte(`{
		"Id": 1,
		"Discription": "My aaaaaa",
		"Title": "My title"
	  }`)

	req, err := http.NewRequest("POST", "/Thread", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Id Already Exists`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestInvalidPostThread(t *testing.T) {

	var jsonStr = []byte(`{
		"Id": "1a",
		"Discription": "My aaaaaa",
		"Title": "My title"
	  }`)

	req, err := http.NewRequest("POST", "/Thread", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `There was an error Processing your request`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPutThread(t *testing.T) {

	var jsonStr = []byte(`{
		"Id": "1",
		"Discription": "My aaaaaa",
		"Title": "My title"
	  }`)

	req, err := http.NewRequest("PUT", "/Thread", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequest)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `There was an error Processing your request`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
