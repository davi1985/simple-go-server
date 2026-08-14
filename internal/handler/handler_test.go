/*
 * Package declaration for testing. The '_test' suffix indicates black-box testing,
 * meaning we test the 'handler' package as an external consumer.
 */
package handler_test

import (
	"go-server/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
 * Test function for the /hello endpoint.
 * The *testing.T parameter provides methods to manage test execution and log failures.
 */
func TestHelloHandler(t *testing.T) {
	/*
	 * http.NewRequest creates an in-memory simulated HTTP request.
	 * Parameters: Method (GET), Route (/hello), Request Body (nil = no body)
	 */
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)

	/*
	 * If request creation fails (e.g., invalid method or URL format),
	 * t.Fatalf stops the test execution immediately and prints the error.
	 */
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	/*
	 * httptest.NewRecorder() creates a recorder implementing http.ResponseWriter.
	 * It captures the status code, response headers, and body sent back by the handler.
	 */
	rr := httptest.NewRecorder()

	// Call the handler function directly, passing the recorder and the mock request.
	handler.Hello(rr, req)

	/*
	 * rr.Code holds the HTTP status code returned by the handler.
	 * We check if it matches 200 OK (http.StatusOK).
	 */
	if status := rr.Code; status != http.StatusOK {
		// t.Errorf logs a test failure message but DOES NOT stop remaining code execution.
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Define the exact body response string expected from the handler.
	expected := "hello!"

	// rr.Body.String() converts the captured response body buffer into a string.
	if rr.Body.String() != expected {
		// Prints the actual value received (got) versus expected value (want) if they differ.
		t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

/*
 * Table-driven test for the /form endpoint.
 * This is the most idiomatic pattern in Go to test multiple scenarios while reusing test logic.
 */
func TestFormHandler(t *testing.T) {
	// Define an anonymous struct slice containing our test cases.
	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		// Case 1: Validating handler behavior for a GET request
		{
			name:           "Valid GET request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		// Case 2: Validating handler behavior for a POST request
		{
			name:           "Valid POST request",
			method:         http.MethodPost,
			expectedStatus: http.StatusOK,
		},
	}

	// Iterate through the test suite using 'range'
	for _, tt := range tests {
		/*
		 * t.Run executes a sub-test for each case in the table.
		 * It enables running and identifying individual test cases in the terminal.
		 */
		t.Run(tt.name, func(t *testing.T) {
			// Create the HTTP request using the HTTP method defined for the current iteration (tt.method)
			req, err := http.NewRequest(tt.method, "/form", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Instantiate a fresh response recorder for this sub-test
			rr := httptest.NewRecorder()

			// Execute the handler with current scenario data
			handler.Form(rr, req)

			// Assert if the returned status matches the expected status for this case
			if rr.Code != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code for method %s: got %v want %v",
					tt.method, rr.Code, tt.expectedStatus)
			}
		})
	}
}
