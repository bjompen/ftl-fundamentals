package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	// t.Parallel()
	type testCase struct {
		a, b, want float64
		name       string
		// want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "2 plus 2"},
		{a: 3, b: 2, want: 5, name: "3 plus 2"},
		{a: 3, b: 3, want: 6, name: "3 plus 3"},
	}
	for _, tc := range testCases {
		t.Logf("testCase: %s. testing %f , %f : expecting %f", tc.name, tc.a, tc.b, tc.want)
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	// t.Parallel()
	type testCase struct {
		a, b, want float64
		name       string
	}

	testCases := []testCase{
		{a: 3, b: 2, want: 1, name: "3 minus 2"},
		{a: 4, b: 2, want: 2, name: "4 minus 2"},
		{a: 3, b: 4, want: -1, name: "3 minus 4"},
	}
	for _, tc := range testCases {
		t.Logf("Testcase: %s, testing %f minus %f, expecting %f", tc.name, tc.a, tc.b, tc.want)
		var want float64 = tc.want
		got := calculator.Subtract(tc.a, tc.b)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	// t.Parallel()
	type testCase struct {
		a, b, want int32
		name       string
	}
	testCases := []testCase{
		{a: 2, b: 3, want: 6, name: "2 times 3"},
		{a: 3, b: 4, want: 12, name: "3 times 4"},
		{a: 1, b: 0, want: 0, name: "1 times 0"},
	}
	for _, tc := range testCases {
		t.Logf("Testcase %s, %d times %d is %d", tc.name, tc.a, tc.b, tc.want)
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("ville ha %02d, men fick %02d", tc.want, got)
		}
	}
}

func TestDivision(t *testing.T) {
	// t.Parallel()
	type testCase struct {
		a, b, want  float64
		name        string
		errExpected bool
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2, name: "4 divided by 2", errExpected: false},
		{a: 3, b: 2, want: 1.5, name: "3 divided by 2", errExpected: false},
		{a: 1, b: 0, want: 0, name: "1 divided by 0", errExpected: true},
	}
	for _, tc := range testCases {
		got, er := calculator.Divide(tc.a, tc.b)

		errReceived := er != nil
		if tc.errExpected != errReceived {
			t.Logf("Testcase %s, expected fail", tc.name)
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Logf("Testcase %s, %f divided by %f is %f", tc.name, tc.a, tc.b, tc.want)
			t.Errorf("ville ha %f, men fick %f", tc.want, got)
		}
	}
}

func TestRand(t *testing.T) {
	for i := 1; i < 101; i++ {
		a := rand.Float64()
		b := rand.Float64()

		rndErr := rand.Intn(2)
		if rndErr == 1 {
			want := a + b + 1
			got := calculator.Add(a, b)
			t.Logf("random test %d, including error", i)
			if want != got {
				t.Errorf("matte är svårt! %f + %f = %f, ville ha %f", a, b, got, want)
			}
		}
		if rndErr == 0 {
			want := a + b
			got := calculator.Add(a, b)
			t.Logf("random test %d", i)
			if want != got {
				t.Errorf("matte är svårt! %f + %f = %f, ville ha %f", a, b, got, want)
			}
		}
	}
}
