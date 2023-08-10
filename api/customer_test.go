package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadinessCheck(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	expectedResponseBody := `{"data":"API is running"}`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())

}

func TestGetCustomerInfo(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees/10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	expectedKeys := []string{"id", "name", "age"}

	for _, key := range expectedKeys {
		_, found := responseData[key]
		assert.True(t, found, "Expected key %s not found", key)
	}
}

func TestTestGetCustomerInfo_NoRecord(t *testing.T) {
	record_id := "99999"
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/employees/"+record_id, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	assert.Contains(t, responseData, "error")
	assert.Equal(t, "Cannot find customer information", responseData["error"])

}

func TestUpdateCustomerInfo(t *testing.T) {
	record_id := "4"
	record_id_float, _ := strconv.ParseFloat(record_id, 64)
	router := setupRouter()

	w := httptest.NewRecorder()

	payload := []byte(`{
		"name" : "updated name",
		"age" : 99
	}`)

	req, _ := http.NewRequest("PUT", "/customers/"+record_id, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	assert.Equal(t, record_id_float, responseData["id"])
	assert.Equal(t, "updated name", responseData["name"])
	assert.Equal(t, 99.0, responseData["age"])

}

func TestUpdateCustomerInfo_InvalidRecordID(t *testing.T) {
	record_id := "99999"
	router := setupRouter()

	w := httptest.NewRecorder()

	payload := []byte(`{
		"name" : "updated name",
		"age" : 99
	}`)

	req, _ := http.NewRequest("PUT", "/customers/"+record_id, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	assert.Contains(t, responseData, "error")
	assert.Equal(t, "Cannot find customer information", responseData["error"])
}

func TestUpdateCustomerInfo_InvalidPayloadData(t *testing.T) {
	record_id := "4"
	router := setupRouter()

	w := httptest.NewRecorder()

	// Invalid JSON payload, missing closing curly brace
	payload := []byte(`{
		"name": 3452123,
		"age" : -66eiei
	}`)

	req, _ := http.NewRequest("PUT", "/customers/"+record_id, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	assert.Contains(t, responseData, "error")
	assert.Equal(t, "Invalid customer data", responseData["error"])
}

func TestCreateNewCustomer(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	payload := []byte(`{
		"name" : "test name",
		"age" : 26
	}`)

	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	idValue := responseData["id"]
	assert.IsType(t, float64(0), idValue, "Expected key 'id' to be of type uint")

	assert.Equal(t, "test name", responseData["name"])
	assert.Equal(t, 26.0, responseData["age"])

}

func TestDeleteCustomer(t *testing.T) {
	record_id := "14" //or some existing id
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/employees/"+record_id, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	idValue := responseData["id"]
	assert.IsType(t, float64(11), idValue, "Expected key 'id' to be of type uint")

}

func TestDeleteCustomer_NoRecord(t *testing.T) {
	record_id := "99999"
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/employees/"+record_id, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var responseData map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseData)
	assert.NoError(t, err)

	assert.Contains(t, responseData, "error")
	assert.Equal(t, "Cannot find customer information", responseData["error"])

}
