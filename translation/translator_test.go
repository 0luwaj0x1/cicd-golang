package translation

import (
	"testing"
)
func TestTranslate(t *testing.T) {
   testCases := []struct{
		Word string
		Language string
		Translation string
   }{
	{
		Word: "hello",
		Language: "english",
		Translation: "hello",
	},
	{
		Word: "hello",
		Language: "german",
		Translation: "hallo",
	},
	{
		Word: "hello",
		Language: "finnish",
		Translation: "hei",
	},
	{
		Word: "hello",
		Language: "dutch",
		Translation: "",
	},
   }

   for _, tc := range testCases {
	got := Translate(tc.Word, tc.Language)

	if got != tc.Translation {
		t.Errorf("want %s to be %s from %s but got %s", tc.Word, tc.Language, tc.Translation, got)
	}
   }
}