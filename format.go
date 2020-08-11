package expr

import (
	"fmt"
	"strings"
)

func format(tokens []token, isConcat bool) (string, int) {
	var sb strings.Builder
	if isConcat {
		sb.WriteString("(")
	}
	j := -1
	concatCallback := false
	for i, t := range tokens {
		if i <= j {
			continue
		}
		j++
		switch t.kind {
		case name:
			if len(tokens) > i+1 && tokens[i+1].is(punctuation) && tokens[i+1].value == "(" {
				//function
				f, pos := funcTranscode(t.value)
				if f == "+" {
					concatCallback = true
				} else {
					concatCallback = false
				}
				if pos == -1 {
					return "", j
				}
				if pos == 0 {
					sb.WriteString(f)
					sb.WriteString("(")
					str, p := format(tokens[i+1:], concatCallback)
					sb.WriteString(str)
					j += p
					sb.WriteString(")")
				} else if pos == 1 {
					sb.WriteString("(")
					str, p := format(tokens[i+1:], concatCallback)
					sb.WriteString(str)
					j += p
					sb.WriteString(")")
				} else if pos == 2 {
					sb.WriteString("(")
					str, p := format(tokens[i+1:], concatCallback)
					sb.WriteString(str)
					j += p
					sb.WriteString(")")
					sb.WriteString(f)
				}
			} else {
				sb.WriteString(fmt.Sprintf("doc['%s'].value", t.value))
			}
			break
		case punctuation:
			if t.value == ")" {
				if isConcat {
					sb.WriteString(".toString()")
				}
				return sb.String(), j + 1
			} else if t.value == "," && isConcat {
				sb.WriteString(").toString()+")
			}
			break
		default:
			sb.WriteString(t.value)
		}
	}
	if isConcat {
		sb.WriteString(")")
	}
	return sb.String(), j
}

func funcTranscode(function string) (string, int) {
	//https://www.elastic.co/guide/en/elasticsearch/painless/6.7/painless-api-reference.html
	switch strings.ToLower(function) {
	// DATE  org.joda.time.ReadableDateTime
	case "day_diff":
		//https://stackoverflow.com/questions/53178425/painless-get-time-difference
		return "", -1
	case "day":
		return ".getDayOfYear()", 2
	case "hour":
		return ".getHourOfDay()", 2
	case "minute":
		return "getMinuteOfHour()", 2
	case "second":
		return ".getSecondOfMinute()", 2
	case "month":
		return ".getMonthOfYear()", 2
	case "quarter":
		return "", -1
	case "week":
		return ".getWeekyear()", 2
	case "weekday":
		return ".getDayOfWeek()", 2
	case "year":
		return ".getYear()", 2
		// MATH
	case "abs":
		return "StrictMath.abs", 0
	case "acos":
		return "StrictMath.acos", 0
	case "asin":
		return "StrictMath.asin", 0
	case "atan":
		// SATAN !!!
		return "StrictMath.atan", 0
	case "ceil":
		return "StrictMath.ceil", 0
	case "cos":
		return "StrictMath.cos", 0
	case "floor":
		return "StrictMath.floor", 0
	case "log":
		return "StrictMath.log", 0
	case "log10":
		return "StrictMath.log10", 0
	case "pow":
		return "StrictMath.pow", 0
	case "round":
		return "StrictMath.round", 0
	case "sin":
		return "StrictMath.sin", 0
	case "tan":
		return "StrictMath.tan", 0
	// TEXT
	case "upper":
		return ".toUpperCase()", 2
	case "lower":
		return ".toLowerCase()", 2
	case "length":
		return ".length()", 2
	case "trim":
		return ".trim()", 2
		// TODO : A revoir
	case "concat":
		return "+", 1
	default:
		return "", -1
	}
}
