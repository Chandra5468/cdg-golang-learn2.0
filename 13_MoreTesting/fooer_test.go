package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(3)

	if result != "Foo" {
		t.Errorf("Result was incorrect, got %s, want %s", result, "Foo")
	}
}

// Table driven tests

func TestFooerTableDriven(t *testing.T) {
	// Defining the columns of the table

	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be Foo", 9, "Foo"},
		{"5 should not be Foo", 5, "5"},
		{"0 should be Foo", 0, "Foo"},
		{"3 should not be Foo", 3, "Foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fooer(tt.input)

			if result != tt.want {
				t.Errorf("For %s the input was %d but we got %s", tt.name, tt.input, tt.want)
			}
		})
	}
}

//----------------------------------------------------------------------------------------
// Parallel Testing : By default, tests are run sequentially;
// the method Parallel() signals that a test should be run in parallel.

//go test handles parallel tests by pausing each test that calls t.Parallel(),
// and then resuming them in parallel when all non-parallel tests have been completed.
// The GOMAXPROCS environment defines how many tests can run in parallel at one time,
// and by default this number is equal to the number of CPUs.

func TestFooersParallel(t *testing.T) {
	t.Run("Test 3 running in parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was expecting Foo but got %s", result)
		}
	})

	t.Run("Test 7 running in parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(9)

		if result != "Foo" {
			t.Errorf("Result was expecting Foo but got %s", result)
		}
	})
}

// Benchmark Testing :
func BenchmarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fooer(i)
	}
}

// go test -bench=. or go test -bench=Fooer

/// FUZZ TESTS ::::
func FuzzFooer(f *testing.F) {
	f.Add(100)

	f.Fuzz(func(t *testing.T, a int) {
		Fooer(a)
	})
}
