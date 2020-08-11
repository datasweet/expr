package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormating(t *testing.T) {
	tokens, err := lex("CONCAT(abs(campagne) / annee, upper(`societe.nom`))")
	assert.NoError(t, err)
	out, _ := format(tokens, false)
	assert.Equal(t, `((StrictMath.abs(doc['campagne'].value)/doc['annee'].value).toString()+(doc['societe.nom'].value).toUpperCase().toString())`, out)
}

func TestFormating2(t *testing.T) {
	tokens, err := lex("upper(`societe.nom`)")
	assert.NoError(t, err)
	out, _ := format(tokens, false)
	assert.Equal(t, `(doc['societe.nom'].value).toUpperCase()`, out)
}

func TestFormating3(t *testing.T) {
	tokens, err := lex("foo")
	assert.NoError(t, err)
	out, _ := format(tokens, false)
	assert.Equal(t, `doc['foo'].value`, out)
}

func TestFormating4(t *testing.T) {
	tokens, err := lex("3 / 4")
	assert.NoError(t, err)
	out, _ := format(tokens, false)
	assert.Equal(t, `3/4`, out)
}

func TestFormating5(t *testing.T) {
	tokens, err := lex("day(timestamp)")
	assert.NoError(t, err)
	out, _ := format(tokens, false)
	assert.Equal(t, `(doc['timestamp'].value).getDayOfYear()`, out)
}

func TestFormating6(t *testing.T) {
	tokens, err := lex("Upper(day(timestamp)*sin(3))")
	assert.NoError(t, err)
	out, _ := format(tokens, false)
	assert.Equal(t, `((doc['timestamp'].value).getDayOfYear()*StrictMath.sin(3)).toUpperCase()`, out)
}
