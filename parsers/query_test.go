package parsers

import (
	"net/url"
	"testing"
)

func TestNormalizeURLQueryString(t *testing.T) {
	tests := []struct {
		input          string
		queryString    bool
		qsIgnoreValues bool
		want           string
	}{
		{input: "/sales", want: "/sales"},
		{input: "/sales?", want: "/sales"},
		{input: "/sales?x=1", want: "/sales"},
		{input: "/sales?", qsIgnoreValues: true, want: "/sales"},
		{input: "/sales?x=1", qsIgnoreValues: true, want: "/sales"},
		{input: "/sales", queryString: true, want: "/sales"},
		{input: "/sales?", queryString: true, want: "/sales?"},
		{input: "/sales?x=1", queryString: true, want: "/sales?x=1"},
		{input: "/sales?b=2&a=1", queryString: true, want: "/sales?a=1&b=2"},
		{input: "/sales?", queryString: true, qsIgnoreValues: true, want: "/sales?"},
		{input: "/sales?x=1", queryString: true, qsIgnoreValues: true, want: "/sales?x=xxx"},
		{input: "/sales?x=2", queryString: true, qsIgnoreValues: true, want: "/sales?x=xxx"},
		{input: "/sales%3F", want: "/sales%3F"},
		{input: "/sales%3F?", want: "/sales%3F"},
		{input: "/sales%3F?x=1", want: "/sales%3F"},
		{input: "/sales%3F?", queryString: true, want: "/sales%3F?"},
		{input: "/sales%3F?x=1", queryString: true, want: "/sales%3F?x=1"},
		{input: "/sales%2Fdetails?", want: "/sales%2Fdetails"},
		{input: "/sales%2Fdetails?x=1", want: "/sales%2Fdetails"},
		{input: "/sales%2Fdetails?", queryString: true, want: "/sales%2Fdetails?"},
		{input: "/sales%2Fdetails?x=1", queryString: true, want: "/sales%2Fdetails?x=1"},
	}

	for _, test := range tests {
		t.Run(test.input+"/"+test.want, func(t *testing.T) {
			src, err := url.Parse(test.input)
			if err != nil {
				t.Fatal(err)
			}
			original := *src

			got := normalizeURL(src, test.queryString, test.qsIgnoreValues)
			if got.String() != test.want {
				t.Fatalf("normalized URL = %q, want %q", got.String(), test.want)
			}
			if !test.queryString && (got.RawQuery != "" || got.ForceQuery) {
				t.Fatalf("query excluded but RawQuery = %q, ForceQuery = %v", got.RawQuery, got.ForceQuery)
			}
			if *src != original {
				t.Fatalf("source URL changed: got %#v, want %#v", *src, original)
			}
		})
	}
}
