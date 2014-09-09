package logger

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestGeneralUsage(t *testing.T) {
	CleanUp()
	Log("holi")
	time.Sleep(2 * time.Second)

	fBytes, _ := ioutil.ReadFile("fortinet.log")
	lines := strings.Split(string(fBytes), "\n")
	matched, _ := regexp.MatchString("\\d{4}/\\d{2}/\\d{2} \\d{2}:\\d{2}:\\d{2} holi", lines[0])

	if !matched {
		t.Fatalf("Where is my line? %s", lines[0])
	}
}

func CleanUp() {
	os.Remove("fortinet.log")
}
