package expr_test

import "testing"

func TestFuncAggAvg(t *testing.T) {
	testEval(t, []evalTest{
		{
			"avg()",
			nil,
			nil,
		},
		{
			"avg(1,2,3,4,5)",
			nil,
			float64(3),
		},
		{
			"avg(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(3),
		},
		{
			"avg(1,2,3,4,'5')",
			nil,
			float64(3),
		},
		{
			"avg(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggCount(t *testing.T) {
	testEval(t, []evalTest{
		{
			"count()",
			nil,
			0,
		},
		{
			"count(1,2,3,4,5)",
			nil,
			5,
		},
		{
			"count(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			5,
		},
		{
			"count(1,2,true,4,'5')",
			nil,
			5,
		},
	})
}

func TestFuncAggCountDistinct(t *testing.T) {
	testEval(t, []evalTest{
		{
			"count_distinct()",
			nil,
			0,
		},
		{
			"count_distinct(1,2,2,3,2)",
			nil,
			3,
		},
		{
			"count_distinct(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 2, 3, 2}},
			3,
		},
		{
			"count_distinct(1, 2, '2', 3, 2)",
			nil,
			3,
		},
	})
}

func TestFuncAggCusum(t *testing.T) {
	testEval(t, []evalTest{
		{
			"cusum()",
			nil,
			nil,
		},
		{
			"cusum(1,2,3,4,5)",
			nil,
			[]float64{1, 3, 6, 10, 15},
		},
		{
			"cusum(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			[]float64{1, 3, 6, 10, 15},
		},
		{
			"cusum(1,2,3,4,'5')",
			nil,
			[]float64{1, 3, 6, 10, 15},
		},
		{
			"cusum(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggMax(t *testing.T) {
	testEval(t, []evalTest{
		{
			"max()",
			nil,
			nil,
		},
		{
			"max(1,2,3,4,5)",
			nil,
			float64(5),
		},
		{
			"max(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(5),
		},
		{
			"max(1,2,3,4,'5')",
			nil,
			float64(5),
		},
		{
			"max(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggMedian(t *testing.T) {
	testEval(t, []evalTest{
		{
			"median()",
			nil,
			nil,
		},
		{
			"median(1,2,3,4,5)",
			nil,
			float64(3),
		},
		{
			"median(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(3),
		},
		{
			"median(1,2,3,4,'5')",
			nil,
			float64(3),
		},
		{
			"median(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggMin(t *testing.T) {
	testEval(t, []evalTest{
		{
			"min()",
			nil,
			nil,
		},
		{
			"min(1,2,3,4,5)",
			nil,
			float64(1),
		},
		{
			"min(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(1),
		},
		{
			"min(1,2,3,4,'5')",
			nil,
			float64(1),
		},
		{
			"min(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggStddev(t *testing.T) {
	testEval(t, []evalTest{
		{
			"stddev()",
			nil,
			nil,
		},
		{
			"stddev(1,2,3,4,5)",
			nil,
			float64(1.4142135623730951),
		},
		{
			"stddev(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(1.4142135623730951),
		},
		{
			"stddev(1,2,3,4,'5')",
			nil,
			float64(1.4142135623730951),
		},
		{
			"stddev(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggSum(t *testing.T) {
	testEval(t, []evalTest{
		{
			"sum()",
			nil,
			nil,
		},
		{
			"sum(1,2,3,4,5)",
			nil,
			float64(15),
		},
		{
			"sum(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(15),
		},
		{
			"sum(1,2,3,4,'5')",
			nil,
			float64(15),
		},
		{
			"sum(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}

func TestFuncAggVariance(t *testing.T) {
	testEval(t, []evalTest{
		{
			"variance()",
			nil,
			nil,
		},
		{
			"variance(1,2,3,4,5)",
			nil,
			float64(2),
		},
		{
			"variance(foo)",
			map[string]interface{}{"foo": []interface{}{1, 2, 3, 4, 5}},
			float64(2),
		},
		{
			"variance(1,2,3,4,'5')",
			nil,
			float64(2),
		},
		{
			"variance(1,2,3,4,'no')",
			nil,
			nil,
		},
	})
}
