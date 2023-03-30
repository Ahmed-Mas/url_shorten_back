package urlshorten

import (
	"fmt"
	"testing"
)

func testCleanHelper(longURL, expected string) error {
	holder := UrlHolder{
		LongURL: longURL,
	}
	err := holder.cleanURL()
	if err != nil {
		return fmt.Errorf("couldnt validate url: %v", err)
	}

	if holder.LongURL != expected {
		return fmt.Errorf("url not clean:\n\texpected:%s\n\tgot:%s", expected, holder.LongURL)
	}
	return nil
}

func TestCleanURL(t *testing.T) {
	longURL := "http://example.com"
	expected := "http://example.com"
	err := testCleanHelper(longURL, expected)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateURL2(t *testing.T) {
	longURL := "https://google.com"
	expected := "https://google.com"
	err := testCleanHelper(longURL, expected)
	if err != nil {
		t.Error(err)
	}
}

func TestCleanURL3(t *testing.T) {
	longURL := "google.com"
	expected := "https://google.com"
	err := testCleanHelper(longURL, expected)
	if err != nil {
		t.Error(err)
	}
}

func TestCleanURLFail(t *testing.T) {
	longURL := "random"
	expected := "https://random"
	err := testCleanHelper(longURL, expected)
	if err != nil {
		t.Error(err)
	}
}
