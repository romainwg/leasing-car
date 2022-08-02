package api

import "testing"

func TestGetDataBaseURL(t *testing.T) {
	var expected string = "postgres://abc:def@ghi:jkl/mno"
	ans := GetDataBaseURL("abc", "def", "ghi", "jkl", "mno")
	if ans != expected {
		t.Errorf("TestGetDataBaseURL(\"abc\", \"def\", \"ghi\", \"jkl\", \"mno\") = %s; want %s", ans, expected)
	}
}