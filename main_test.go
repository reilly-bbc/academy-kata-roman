package main

import (
	"testing"
)

func TestValidNumerals(t *testing.T) {
	testCases := []struct {
		roman  string
		arabic int
	}{
		{
			"I", 1,
		},
		{
			"II", 2,
		},
		{
			"III", 3,
		},
		{
			"IV", 4,
		},
		{
			"V", 5,
		},
		{
			"VI", 6,
		},
		{
			"VII", 7,
		},
		{
			"VIII", 8,
		},
		{
			"IX", 9,
		},
		{
			"X", 10,
		},
		{
			"XI", 11,
		},
		{
			"XX", 20,
		},
		{
			"XXX", 30,
		},
		{
			"XL", 40,
		},
		{
			"L", 50,
		},
		{
			"C", 100,
		},
		{
			"CC", 200,
		},
		{
			"CCC", 300,
		},
		{
			"CD", 400,
		},
		{
			"D", 500,
		},
		{
			"CM", 900,
		},
		{
			"M", 1000,
		},
		{
			"MM", 2000,
		},
		{
			"MMM", 3000,
		},
		{
			"MCDLVXI", 1456,
		},
		{
			"MCMXCIX", 1999,
		},
		{
			"MMMCMXCIX", 3999,
		},
	}

	for _, testCase := range testCases {
		result, err := romanToArabic(testCase.roman)
		if err != nil {
			t.Errorf("received error for %s", testCase.roman)
		}
		if result != testCase.arabic {
			t.Errorf("expected '%d', got '%d' for %s", testCase.arabic, result, testCase.roman)
		}
	}
}

func TestInvalidNumerals(t *testing.T) {
	tooManyUnitsError := "more than three sequential units used"
	tooManyFivesError := "more than one sequential fives used"
	invalidCharacterError := "invalid character in provided numerals"
	badSubtractionError := "invalid subtraction in sequence"
	nonDecreasingSeqError := "sequence is non-decreasing"

	testCases := []struct {
		roman string
		err   string
	}{
		{
			"IIII", tooManyUnitsError,
		},
		{
			"MCCCCVI", tooManyUnitsError,
		},
		{
			"VV", tooManyFivesError,
		},
		{
			"XLLI", tooManyFivesError,
		},
		{
			"ABCDE", invalidCharacterError,
		},
		{
			"MMMCMXCIQ", invalidCharacterError,
		},
		{
			"IL", badSubtractionError,
		},
		{
			"MIMXCIX", badSubtractionError,
		},
		{
			"MCMIIX", nonDecreasingSeqError,
		},
		{
			"IVXLCDM", nonDecreasingSeqError,
		},
	}
	for _, testCase := range testCases {
		_, err := romanToArabic(testCase.roman)
		if err.Error() != testCase.err {
			t.Errorf("expected %v, got %v", testCase.err, err)
		}
	}
}
