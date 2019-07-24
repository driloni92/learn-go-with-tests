package main

import "testing"

func TestSayHello(t *testing.T) {
	cases := []struct {
		Description string
		Name        string
		Language    string
		Want        string
	}{
		{Description: "Default to English", Name: "Drilon", Language: "unknown-lang", Want: "Hello, Drilon"},
		{Description: "In French", Name: "John", Language: frenchLanguageCode, Want: "Bonjour, John"},
		{Description: "In Spanish", Name: "Christian", Language: spanishLanguageCode, Want: "Hola, Christian"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := SayHello(test.Name, test.Language)
			if got != test.Want {
				t.Errorf("got %s, want %s", got, test.Want)
			}
		})
	}

}
