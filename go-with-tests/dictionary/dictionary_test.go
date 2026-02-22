package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"wow": "very wow"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("wow")
		want := "very wow"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("test")
		assertError(t, err, err)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "wow")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("add", func(t *testing.T) {
		word := "yey"
		def := "very yey"
		dictionary.Add(word, def)

		want := "very yey"
		assertDefinition(t, dictionary, "yey", want)
	})

	t.Run("update", func(t *testing.T) {
		word := "yey"
		def := "updated"
		dictionary.Add(word, def)

		want := "updated"
		assertDefinition(t, dictionary, "yey", want)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}