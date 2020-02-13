package metrics

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// FizzBuzzMetricsSuite tests all the functionality
// that Metrics should implement
func FizzBuzzMetricsSuite(t *testing.T, newMetrics func() Metrics) {
	tests := map[string]struct {
		args []Request
		want []Result
	}{
		"no requests": {
			args: []Request{},
			want: []Result{},
		},
		"one request": {
			args: []Request{
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
			},
			want: []Result{
				{
					Request: Request{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
					Hits:    1,
				},
			},
		},
		"5 similar requests": {
			args: []Request{
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
			},
			want: []Result{
				{
					Request: Request{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
					Hits:    5,
				},
			},
		},
		"3 requests with two similar": {
			args: []Request{
				{Int1: 2, Int2: 4, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
				{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
			},
			want: []Result{
				{
					Request: Request{Int1: 2, Int2: 4, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
					Hits:    1,
				},
				{
					Request: Request{Int1: 3, Int2: 5, Limit: 10, Str1: "Fizz", Str2: "Buzz"},
					Hits:    2,
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := NewInMemoryMetrics()

			for _, req := range tt.args {
				m.Record(req)
			}

			got := m.Get()

			if !cmp.Equal(
				got, tt.want,
				cmpopts.SortSlices(func(x, y Result) bool { return x.Hits < y.Hits }),
			) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

// TestInMemoryMetrics uses the FizzBuzzMetricsSuite to test the
// in memory implementation of the Metrics interface.
func TestInMemoryMetrics(t *testing.T) {
	newMetrics := func() Metrics {
		m := NewInMemoryMetrics()
		return &m
	}

	FizzBuzzMetricsSuite(t, newMetrics)
}
