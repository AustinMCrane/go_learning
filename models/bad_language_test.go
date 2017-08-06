package models

import (
	"fmt"
	"testing"
)

type BadLanguageTest struct {
	s        string
	expected bool
}

func TestHasBadLanguage(t *testing.T) {
	tests := [10]BadLanguageTest{
		BadLanguageTest{"Fuck me good", true},
		BadLanguageTest{"That basterd", true},
		BadLanguageTest{"That shit", true},
		BadLanguageTest{"That raper", true},
		BadLanguageTest{"That asshole", true},
		BadLanguageTest{"Something something asshole", true},
		BadLanguageTest{"Something pussy asshole", true},
		BadLanguageTest{"Something pus sy", false},
		BadLanguageTest{"Something something make it good", false},
		BadLanguageTest{"can has idk make it good", false},
	}
	for i := 0; i < len(tests); i++ {
		test := tests[i]
		expected, badWord := HasBadLanguage(test.s)
		if expected != test.expected {
			t.Error(fmt.Sprintf("Test number %d failed, found %s", i, badWord))
		}
	}
}
