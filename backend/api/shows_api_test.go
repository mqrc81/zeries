package api

import (
	"testing"

	"github.com/nbio/st"
)

func Test_calculateRange(t *testing.T) {
	type args struct {
		pageQueryParam string
		pastReleases   int
	}
	type want struct {
		amount int
		offset int
	}
	tests := []struct {
		name string
		args
		want
	}{
		{
			name: "#1",
			args: args{pageQueryParam: "", pastReleases: 55},
			want: want{amount: 20, offset: 55},
		},
		{
			name: "#2",
			args: args{pageQueryParam: "1", pastReleases: 55},
			want: want{amount: 20, offset: 75},
		},
		{
			name: "#3",
			args: args{pageQueryParam: "-1", pastReleases: 55},
			want: want{amount: 20, offset: 35},
		},
		{
			name: "#4",
			args: args{pageQueryParam: "-3", pastReleases: 55},
			want: want{amount: 15, offset: 0},
		},
		{
			name: "#5",
			args: args{pageQueryParam: "3", pastReleases: 55},
			want: want{amount: 20, offset: 115},
		},
		{
			name: "#6",
			args: args{pageQueryParam: "abc", pastReleases: 55},
			want: want{amount: 20, offset: 55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAmount, gotOffset := calculateRange(tt.args.pageQueryParam, tt.args.pastReleases)

			st.Expect(t, gotAmount, tt.amount)
			st.Expect(t, gotOffset, tt.offset)

		})
	}
}
