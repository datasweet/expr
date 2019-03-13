package expr_test

import "testing"

func TestFuncTextConcat(t *testing.T) {
	testEval(t, []evalTest{
		{
			"'sea' ~ 'food'",
			nil,
			"seafood",
		},
		{
			"'sea' ~ nil",
			nil,
			"sea",
		},
		{
			"'sea' ~ 1",
			nil,
			"sea1",
		},
		{
			"1 ~ 2",
			nil,
			"12",
		},
		{
			"country ~ '_' ~ upper(country)",
			map[string]interface{}{"country": []interface{}{"fr", "de", "es"}},
			[]interface{}{"fr_FR", "de_DE", "es_ES"},
		},
	})
}

func TestFuncTextLength(t *testing.T) {
	testEval(t, []evalTest{
		{
			"length(nil)",
			nil,
			0,
		},
		{
			"length('hello world !')",
			nil,
			13,
		},
		{
			"length(country)",
			map[string]interface{}{"country": []interface{}{"FRANCE", "GERMANY"}},
			[]interface{}{6, 7},
		},
	})
}

func TestFuncTextLower(t *testing.T) {
	testEval(t, []evalTest{
		{
			"lower(nil)",
			nil,
			"",
		},
		{
			"lower('=== RUN   TestFuncTextLower')",
			nil,
			"=== run   testfunctextlower",
		},
		{
			"lower(country)",
			map[string]interface{}{"country": []interface{}{"FrANcE", "GeRMaNY"}},
			[]interface{}{"france", "germany"},
		},
	})
}

func TestFuncTextMatches(t *testing.T) {
	testEval(t, []evalTest{
		{
			`"seafood" matches "foo.*"`,
			nil,
			true,
		},
		{
			`"seafood" matches "sea" ~ "food"`,
			nil,
			true,
		},
		{
			`not ("seafood" matches "[0-9]+") ? "a" : "b"`,
			nil,
			"a",
		},
		{
			"`food` matches \"sea\" ~ \"|\" ~ \"food\"",
			map[string]interface{}{"food": []interface{}{"sea", "is", "more", "food"}},
			[]interface{}{true, false, false, true},
		},
	})
}

func TestFuncTextTrim(t *testing.T) {
	testEval(t, []evalTest{
		{
			"trim(nil)",
			nil,
			"",
		},
		{
			"trim('\n\r               === RUN   TestFuncTextLower   \t\r\n       ')",
			nil,
			"=== RUN   TestFuncTextLower",
		},
		{
			"trim(country)",
			map[string]interface{}{"country": []interface{}{"  FrANcE  ", "      GeRMaNY     "}},
			[]interface{}{"FrANcE", "GeRMaNY"},
		},
	})
}

func TestFuncTextUpper(t *testing.T) {
	testEval(t, []evalTest{
		{
			"upper(nil)",
			nil,
			"",
		},
		{
			"upper('=== RUN   TestFuncTextLower')",
			nil,
			"=== RUN   TESTFUNCTEXTLOWER",
		},
		{
			"upper(country)",
			map[string]interface{}{"country": []interface{}{"FrANcE", "GeRMaNY"}},
			[]interface{}{"FRANCE", "GERMANY"},
		},
	})
}
