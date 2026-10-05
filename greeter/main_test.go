package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	cfg := loadConfig()
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Alice", nil)
	rec := httptest.NewRecorder()

	newMux(cfg).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if want := "Hello, Alice!"; got.Message != want {
		t.Errorf("message = %q, want %q", got.Message, want)
	}
}

func TestHelloWithoutName(t *testing.T) {
	cfg := loadConfig()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	newMux(cfg).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Message == "" {
		t.Error("message is empty, want a default greeting")
	}
}

func TestHelloMethodNotAllowed(t *testing.T) {
	cfg := loadConfig()
	req := httptest.NewRequest(http.MethodPost, "/hello", nil)
	rec := httptest.NewRecorder()

	newMux(cfg).ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Error == "" {
		t.Error("error message is empty")
	}
}

func TestFarewellWithName(t *testing.T) {
	cfg := loadConfig()
	req := httptest.NewRequest(http.MethodGet, "/farewell?name=Alice", nil)
	rec := httptest.NewRecorder()

	newMux(cfg).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if want := "Goodbye, Alice!"; got.Message != want {
		t.Errorf("message = %q, want %q", got.Message, want)
	}
}

func TestFarewellWithoutName(t *testing.T) {
	cfg := loadConfig()
	req := httptest.NewRequest(http.MethodGet, "/farewell", nil)
	rec := httptest.NewRecorder()

	newMux(cfg).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var got Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Message == "" {
		t.Error("message is empty, want a default farewell")
	}
}

func TestFarewellMethodNotAllowed(t *testing.T) {
	cfg := loadConfig()
	req := httptest.NewRequest(http.MethodPost, "/farewell", nil)
	rec := httptest.NewRecorder()

	newMux(cfg).ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got.Error == "" {
		t.Error("error message is empty")
	}
}
