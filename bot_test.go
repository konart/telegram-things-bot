package main

import "testing"

func TestGetTitleBody(t *testing.T) {
	testCases := []struct {
		text string
		expectedTitle string
		expectedBody string
	}{
		{"Title\tBody", "Title", "Body"},
		{"", "", ""},
		{"Some title", "Some title", ""},
	}

	for i, tc := range testCases {
		title, body := getTitleBody(tc.text)
		if title != tc.expectedTitle || body != tc.expectedBody {
			t.Errorf("TestCase #%d: expected title '%s' and body '%s', got" +
				"'%s' and '%s' instead", i, tc.expectedTitle, tc.expectedBody, title, body)
		}
	}
}