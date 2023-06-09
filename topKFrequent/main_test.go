package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 5},
				k:    3,
			},
			want: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topKFrequent([]int{1, 1, 1, 2, 2, 3, 3}, 3)
	}
}

//func BenchmarkTopKFrequent2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		topKFrequent2([]int{1, 1, 1, 2, 2, 3, 3}, 3)
//	}
//}
//
//func BenchmarkTopKFrequent3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		topKFrequent3([]int{1, 1, 1, 2, 2, 3, 3}, 3)
//	}
//}
//
//func BenchmarkTopKFrequent4(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		topKFrequent4([]int{1, 1, 1, 2, 2, 3, 3}, 3)
//	}
//}
