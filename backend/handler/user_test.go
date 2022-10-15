package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SuperTikuwa/book_app/handler"
)

func TestStoreUser(t *testing.T) {
	r := handler.StoreUserRequest{
		Name:  faker.FirstName(),
		Email: faker.Email(),
	}

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	reqBody := bytes.NewBufferString(string(b))
	req := httptest.NewRequest("POST", "/api/user", reqBody)
	res := httptest.NewRecorder()

	handler.StoreUser(res, req)

	if res.Code != http.StatusCreated {
		t.Log(res.Body.String())
		t.Errorf("Status code is wrong. got %v want %v", res.Code, http.StatusCreated)
	}
}

func TestStoreUser_失敗時に削除できるか(t *testing.T) {
	r := handler.StoreUserRequest{
		Name:  "test",
		Email: faker.Email(),
	}

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	reqBody := bytes.NewBufferString(string(b))
	req := httptest.NewRequest("POST", "/api/user", reqBody)
	res := httptest.NewRecorder()

	handler.StoreUser(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Log(res.Body.String())
		t.Errorf("Status code is wrong. got %v want %v", res.Code, http.StatusCreated)
	}
}
