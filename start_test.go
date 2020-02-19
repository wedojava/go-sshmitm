package main

import (
	"os"
)

func ExampleStart() {
	os.Args = []string{"start"}
	main()
	// Output: Fuck u jiu zai tomorrow!
}

func ExampleStart1()  {
	os.Args = []string{"start", "go"}
	main()
	// Output: starting ssh server on port 22 ...
}

func ExampleStart2()  {
	os.Args = []string{"start", "go", "2222"}
	main()
	// Output: starting ssh server on port 2222 ...
}

func ExampleStart3()  {
	os.Args = []string{"start", "go", "1234", "./test"}
	main()
	// Output: starting ssh server on port 1234 ...
}

//func TestStart(t *testing.T) {
//	tcs := []struct {
//		args   []string
//		except []string
//	}{
//		{[]string{"start"}, []string{"Fuck u jiu zai tomorrow!"}},
//		{[]string{"start", "go"}, []string{"22", "./id_rsa"}},
//		{[]string{"start", "go", "2222"}, []string{"2222", "./id_rsa"}},
//		{[]string{"start", "go", "1234", "./test"}, []string{"1234", "./test"}},
//	}
//
//	for _, tc := range tcs {
//		os.Args = tc.args
//		out = new(bytes.Buffer)
//		main()
//		got := out.(*bytes.Buffer).String()
//		for _, arg := range tc.except {
//			if !strings.Contains(got, arg) {
//				t.Errorf("Err: Got: %s, Except contain %s", got, arg)
//			}
//		}
//	}
//}
