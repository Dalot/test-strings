package main

import (
	"reflect"
	"testing"
)

// Estimate: 5mins
// Real: 5mins
func Test_testValidity(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "success",
			str:  "23-ab-48-caba-56-haha",
			want: true,
		},
		{
			name: "success2",
			str:  "23",
			want: true,
		},
		{
			name: "success2",
			str:  generate(true),
			want: true,
		},
		{
			name: "fail",
			str:  "23-ab-4u-caba-56-haha",
			want: false,
		},
		{
			name: "fail2",
			str:  "23-ab-48-caba-56-h4ha",
			want: false,
		},
		{
			name: "fail3",
			str:  "23-ab-48-caba-56-haha-",
			want: false,
		},
		{
			name: "fail4",
			str:  "-23-ab-48-caba-56-haha",
			want: false,
		},
		{
			name: "fail5",
			str:  generate(false),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.str); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Estimate: 1mins
// Real: 1mins
func Test_averageNumber(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "success",
			str:  "23-ab-48-caba-56-haha",
			want: 4,
		},
		{
			name: "success2",
			str:  "22",
			want: 2,
		},
		{
			name: "success3",
			str:  "1",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageNumber(tt.str); got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Estimate: <5mins
// Real: <5mins
func Test_wholeStory(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "success",
			str:  "23-hello-55-world",
			want: "hello world",
		},
		{
			name: "success2",
			str:  "23-hello",
			want: "hello",
		},
		{
			name: "success",
			str:  "23-hello-99-world-333",
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wholeStory(tt.str); got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Estimate:5mins
// Real: 10mins
func Test_storyStats(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want *Stats
	}{
		{
			name: "success",
			str:  "23-hello-99-worldzz-333-aaabbb",
			want: &Stats{
				shortestWord:      "hello",
				longestWord:       "worldzz",
				avgWordLength:     6,
				avgWordLengthList: []string{"aaabbb"},
			},
		},
		{
			name: "success2",
			str:  "23-hello-99-worldzz-333",
			want: &Stats{
				shortestWord:  "hello",
				longestWord:   "worldzz",
				avgWordLength: 6,
			},
		},
		{
			name: "success3",
			str:  "23-hello-99-world-333",
			want: &Stats{
				shortestWord:      "hello",
				longestWord:       "hello",
				avgWordLength:     5,
				avgWordLengthList: []string{"hello", "world"},
			},
		},
		{
			name: "success4",
			str:  "23",
			want: &Stats{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storyStats(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storyStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
