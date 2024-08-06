package emailsender

import (
	"os"
	"testing"
)

func TestGetEmailBody(t *testing.T) {

	tmpFile, err := os.CreateTemp("", "test_template_*.html")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	defer func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			t.Errorf("Failed to remove temporary file: %v", err)
		}
	}()

	testTemplate := `<html><body>Hello {{.Name}}, your task is {{.AssignTask}}</body></html>`
	if _, err := tmpFile.Write([]byte(testTemplate)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		t.Errorf("Failed to close temporary file: %v", err)
	}

	body, err := getEmailBody(tmpFile.Name(), "TestUser", "TestTask")
	if err != nil {
		t.Fatalf("getEmailBody failed: %v", err)
	}

	expected := "<html><body>Hello TestUser, your task is TestTask</body></html>"
	if string(body) != expected {
		t.Errorf("getEmailBody returned unexpected body. Got %s, want %s", string(body), expected)
	}
}
