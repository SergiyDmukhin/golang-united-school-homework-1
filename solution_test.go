package solution

import (
	"strings"
	"testing"
	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	expected := emoji.Sprint("Hello :map:!")
	msg := GetMessage()

	if !strings.EqualFold(msg, expected) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}
}
