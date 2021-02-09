package main

// func TestCheckRandom(t *testing.T) {
// 	var tests = []struct {
// 		randomStr string
// 		random    int
// 		outLogOut string
// 	}{
// 		{"2", 2, ""},
// 		{"false", 0, ""},
// 		{"true", -1, ""},
// 		{"-10", 0, "random should be a positive integer"},
// 		{"1e1", 0, "random should be either \"false\", \"true\", \"min\", \"max\", \"minMax\", or a positive integer"},
// 	}
// 	for _, test := range tests {
// 		random, logOut := checkRandom(test.randomStr)
// 		if random != test.random {
// 			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.randomStr, test.random, random)
// 		}
// 		if logOut != test.outLogOut {
// 			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.randomStr, test.outLogOut, logOut)
// 		}
// 	}
// }
