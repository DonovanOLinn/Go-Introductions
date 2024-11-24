package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Donovan"},
			result: "Hello, Donovan!",
		},
		{
			items: []string{"Donovan", "Broderick"},
			result: "Hello, Donovan, Broderick!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
	// want is the expected result
	// want := "Hello, test!"
	// got is the actual result
	// got := Say([]string{"test"})

	// if want != got {
	// 	t.Errorf("wanted %s, got %s", want, got)
	// }

	// %v wants it to be printed off in a variable format 
}