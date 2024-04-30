package iteration

import "testing"


// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a")
// 	}
// }

// func ExampleRepeat() {
// 	repeated := Repeat("a")
// 	expected := "aaaaaaa"
// 	//fmt.Println(repeated, expected)
// 	// Output: aaaaaa
// }

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}