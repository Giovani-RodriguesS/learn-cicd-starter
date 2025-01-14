package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	// arrange
	valueWant := ""
	
	// act
	value, _ := GetAPIKey(http.Header{})

	// assert
	if !reflect.DeepEqual(value, valueWant) {
		t.Fatalf("expected: %v, got: %v", valueWant, value)
	}

}
