package main

import "testing"

func TestCleanInput(t *testing.T) {
    
	cases := []struct {
    	input    string
	    expected []string
    }{
	    {
		    input:    "  hello  world  ",
		    expected: []string{"hello", "world"},
	    },
	    {
		    input:    "My name is Gary",
		    expected: []string{"my", "name", "is", "gary"},
	    },
	    {
		    input:    "  BWA,HAHA  ",
		    expected: []string{"bwa,haha"},
	    },
	    {
		    input:    "    ",
		    expected: []string{},
	    },
    }

	for _, c := range cases {
    	actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch\nexpected: %s\nactual: %s", c.expected, actual)
			break
		}
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
    	for i := range actual {
	    	word := actual[i]
	    	expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		    if word != expectedWord {
				t.Errorf("expected: %s\nactual: %s", c.expected, actual)
			    break
			}
    	}
    }
}