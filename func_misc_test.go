package expr_test

import "testing"

func TestFuncMiscContains(t *testing.T) {
	testEval(t, []evalTest{
		{
			"0 in nil",
			nil,
			false,
		},
		{
			"0 not in nil",
			nil,
			true,
		},
		{
			"0 in [1, 2]",
			nil,
			false,
		},
		{
			"0 not in [1, 2]",
			nil,
			true,
		},
		{
			"2 in [1, 2]",
			nil,
			true,
		},
		{
			"2 not in [1, 2]",
			nil,
			false,
		},
		{
			"country in [\"us\", \"fr\"]",
			map[string]interface{}{"country": []interface{}{"fr", "de", "es", "gb"}},
			[]interface{}{true, false, false, false},
		},
		{
			"country not in [\"us\", \"fr\"]",
			map[string]interface{}{"country": []interface{}{"fr", "de", "es", "gb"}},
			[]interface{}{false, true, true, true},
		},
	})
}

func TestFuncMiscCond(t *testing.T) {
	testEval(t, []evalTest{
		{
			`cond > 0 ? "yes" : "no"`,
			map[string]interface{}{"cond": 10, "country": []interface{}{"fr", "de", "es", "gb"}},
			"yes",
		},
		{
			"`cond` > 0 ? `country` : 'no'",
			map[string]interface{}{"cond": []interface{}{10, -20, 30}, "country": []interface{}{"fr", "de", "es", "gb"}},
			[]interface{}{"fr", "no", "es"},
		},
		{
			"`cond` > 0 ? ['fr', 'de', 'es', 'gb'] : 'no'",
			map[string]interface{}{"cond": []interface{}{10, -20, 30}},
			[]interface{}{"fr", "no", "es"},
		},
	})
}
