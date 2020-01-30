package releasetest

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	r := New(OptionName("test_name"), OptionAge(25))
	t.Logf("%+v\n", r)
}
