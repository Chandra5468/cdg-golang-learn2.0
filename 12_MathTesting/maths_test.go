package mathtesting

import "testing"

func TestAdd(t *testing.T) {
	got := add(3, 6)
	want := 9

	if got != want {
		t.Errorf("got %d expected %d", got, want)
	}
}

/* Writing table driven tests

A table-driven test is like a basic unit test except that it maintains a table of different values and results.

*/

type addTest struct {
	testName             string
	arg1, arg2, expected int
}

var addTestCases = []addTest{
	{"addition of positives", 2, 77, 79},
	{"addition of positive and negative", 74, -93, -19},
	{"addition of zeros with negative", 0, -4, -4},
}

func TestAddAll(t *testing.T) {
	for _, test := range addTestCases {
		// You can use t.Run
		// t.Run(test.testName, func(t *testing.T){
		// if result != tt.expected {
		// 	t.Errorf("Add(%d,%d) returned %d but expected %d", tt.a, tt.b, result, tt.expected)
		// }
		//})
		// or simply
		if output := add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("for %s the output %d, is not matching expected %d for arguments %d, %d", test.testName, output, test.expected, test.arg1, test.arg2)
		}
	}
}

/*
	Coverage Tests in GO :
		When writing tests, it is often important to know how much of your actual code the tests cover.
		This is generally referred to as coverage.
		This is also why you have not written a test for your Subtract function in maths.go


	Below command mentions Coverage of Tests in GO :

	in Terminal :
			go test -coverprofile=coverage

			This creates a file called coverage in the same folder of tests and normal package functions.

			PASS
			coverage: 50.0% of statements
			ok      maths   2.552s

	IMP : After go test -coverprofile=coverage

		coverage is a file created

		use below Command :

			go tool cover -html=coverage

			This will redirect you to webpage where what functions are covered and what are not covered is mentioned
*/

/*

	Writing BenchMarks in GO :

	Benchmarking measures the performance of a function or program.
	This allows you to compare implementations and to understand the impact of the changes you make to your code.
	Using that information, you can reveal parts of your Go source code worth optimizing.


	In Go, functions that take the form func BenchmarkXxx(*testing.B) are considered benchmarks.
	go test will execute these benchmarks when you provide the -bench flag.

*/

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(4, 6)
	}
}

// The benchmark function must run the target code b.N times, where N is an integer that can be adjusted

// The --bench flag accepts its arguments in the form of a regular expression.

//**** COMMAND for benchmark testing :
// go test -bench=.
//OR in our case go test -bench=add  (If it is only one function)

// OUTPUT : will be something like this ::
// The resulting output means that the loop ran 10,000,000 times at a speed of 1.07 nanosecond per loop.

// Note: Try not to benchmark your Go code on a busy system being used for other purposes,
// or else you will interfere with the benchmarking process and get inaccurate results
