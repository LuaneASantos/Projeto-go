package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllSuper(t *testing.T) {

	req, err := http.NewRequest("GET", "/all", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	all(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	fmt.Print("teste")
	fmt.Print(len(b) == 0)

	if len(b) == 0 {
		t.Errorf("Esperava registros!")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestAllHeroes(t *testing.T) {

	req, err := http.NewRequest("GET", "/heroes", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	heroes(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if len(b) == 0 {
		t.Errorf("Esperava registros!")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestAllVillains(t *testing.T) {

	req, err := http.NewRequest("GET", "/villains", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	villains(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if len(b) == 0 {
		t.Errorf("Esperava registros!")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestSearchName(t *testing.T) {

	req, err := http.NewRequest("GET", "/search/batman", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	searchName(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if len(b) == 0 {
		t.Errorf("Esperava registros!")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/id/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	byID(rec, req)
	fmt.Print("oi")
	res := rec.Result()

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if len(b) == 0 {
		t.Errorf("Esperava registros!")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}

func TestDeleteByID(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/id/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	removeByID(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	if len(b) != 0 {
		t.Errorf("Erro ao deletar!")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}
