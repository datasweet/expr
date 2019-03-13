package expr_test

import "testing"

func TestFuncDateDateDiff(t *testing.T) {
	testEval(t, []evalTest{
		{
			"date_diff('31/03/2019', '01/03/2019')",
			nil,
			float64(30),
		},
		{
			"date_diff('31/03/2019', '32/03/2019')",
			nil,
			nil,
		},
		{
			"date_diff(nil, nil)",
			nil,
			nil,
		},
	})
}

func TestFuncDateDay(t *testing.T) {
	testEval(t, []evalTest{
		{
			"day(nil)",
			nil,
			nil,
		},
		{
			"day('31/03/2019')",
			nil,
			31,
		},
		{
			"day('32/03/2019')",
			nil,
			nil,
		},
		{
			"day(foo)",
			map[string]interface{}{"foo": []interface{}{"01/03/2019", "15/03/2019", "31/03/2019"}},
			[]interface{}{1, 15, 31},
		},
		{
			"day(foo)",
			map[string]interface{}{"foo": []interface{}{"01/03/2019", "15/03/2019", "32/03/2019"}},
			[]interface{}{1, 15, nil},
		},
	})
}

func TestFuncDateHour(t *testing.T) {
	testEval(t, []evalTest{
		{
			"hour(nil)",
			nil,
			nil,
		},
		{
			"hour('31/03/2019 10:13:47')",
			nil,
			10,
		},
		{
			"hour('31/03/2019')",
			nil,
			0,
		},
		{
			"hour(foo)",
			map[string]interface{}{"foo": []interface{}{"31/03/2019 10:13:47", "31/03/2019"}},
			[]interface{}{10, 0},
		},
		{
			"hour(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 10},
		},
	})
}

func TestFuncDateMinute(t *testing.T) {
	testEval(t, []evalTest{
		{
			"minute(nil)",
			nil,
			nil,
		},
		{
			"minute('31/03/2019 10:13:47')",
			nil,
			13,
		},
		{
			"minute('31/03/2019')",
			nil,
			0,
		},
		{
			"minute(foo)",
			map[string]interface{}{"foo": []interface{}{"31/03/2019 10:13:47", "31/03/2019"}},
			[]interface{}{13, 0},
		},
		{
			"minute(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 13},
		},
	})
}

func TestFuncDateMonth(t *testing.T) {
	testEval(t, []evalTest{
		{
			"month(nil)",
			nil,
			nil,
		},
		{
			"month('31/03/2019 10:13:47')",
			nil,
			3,
		},
		{
			"month('31/03/2019')",
			nil,
			3,
		},
		{
			"month(foo)",
			map[string]interface{}{"foo": []interface{}{"31/03/2019 10:13:47", "31/03/2019"}},
			[]interface{}{3, 3},
		},
		{
			"month(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 3},
		},
	})
}

func TestFuncDateQuarter(t *testing.T) {
	testEval(t, []evalTest{
		{
			"quarter(nil)",
			nil,
			nil,
		},
		{
			"quarter('31/07/2019 10:13:47')",
			nil,
			3,
		},
		{
			"quarter('31/07/2019')",
			nil,
			3,
		},
		{
			"quarter(foo)",
			map[string]interface{}{"foo": []interface{}{"31/07/2019 10:13:47", "31/03/2019"}},
			[]interface{}{3, 1},
		},
		{
			"quarter(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 1},
		},
	})
}

func TestFuncDateSecond(t *testing.T) {
	testEval(t, []evalTest{
		{
			"second(nil)",
			nil,
			nil,
		},
		{
			"second('31/03/2019 10:13:47')",
			nil,
			47,
		},
		{
			"second('31/03/2019')",
			nil,
			0,
		},
		{
			"second(foo)",
			map[string]interface{}{"foo": []interface{}{"31/03/2019 10:13:47", "31/03/2019"}},
			[]interface{}{47, 0},
		},
		{
			"second(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 47},
		},
	})
}

func TestFuncDateWeek(t *testing.T) {
	testEval(t, []evalTest{
		{
			"week(nil)",
			nil,
			nil,
		},
		{
			"week('31/03/2019 10:13:47')",
			nil,
			13,
		},
		{
			"week('31/03/2019')",
			nil,
			13,
		},
		{
			"week(foo)",
			map[string]interface{}{"foo": []interface{}{"31/03/2019 10:13:47", "31/03/2019"}},
			[]interface{}{13, 13},
		},
		{
			"week(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 13},
		},
	})
}

func TestFuncDateYear(t *testing.T) {
	testEval(t, []evalTest{
		{
			"year(nil)",
			nil,
			nil,
		},
		{
			"year('31/03/2019 10:13:47')",
			nil,
			2019,
		},
		{
			"year('31/03/2019')",
			nil,
			2019,
		},
		{
			"year(foo)",
			map[string]interface{}{"foo": []interface{}{"31/03/2019 10:13:47", "31/03/2019"}},
			[]interface{}{2019, 2019},
		},
		{
			"year(foo)",
			map[string]interface{}{"foo": []interface{}{"32/03/2019 10:13:47", "31/03/2019 10:13:47"}},
			[]interface{}{nil, 2019},
		},
	})
}
