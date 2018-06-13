package main

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/KeKsBoTer/sudoku"
)

func Test_readCSV(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		input   string
		want    *sudoku.Field
		wantErr bool
	}{
		{
			name: "valid sudoku file",
			input: `
				7,0,0, 0,0,4, 8,0,0
				0,0,0, 0,0,5, 4,0,0
				0,0,9, 0,0,0, 7,0,0
				4,0,0, 0,0,0, 0,9,0
				8,0,7, 0,0,0, 0,0,0
				0,0,0, 6,1,0, 0,0,0
				0,3,0, 0,5,0, 0,0,1
				0,1,0, 2,0,0, 0,7,5
				0,0,0, 1,4,3, 0,0,0`,
			want: &sudoku.Field{
				{7, 0, 0, 0, 0, 4, 8, 0, 0},
				{0, 0, 0, 0, 0, 5, 4, 0, 0},
				{0, 0, 9, 0, 0, 0, 7, 0, 0},
				{4, 0, 0, 0, 0, 0, 0, 9, 0},
				{8, 0, 7, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 6, 1, 0, 0, 0, 0},
				{0, 3, 0, 0, 5, 0, 0, 0, 1},
				{0, 1, 0, 2, 0, 0, 0, 7, 5},
				{0, 0, 0, 1, 4, 3, 0, 0, 0},
			},
			wantErr: false,
		}, {
			name: "without zeros",
			input: `
				7,,,,,4,8,,
				,,,,,5,4,,
				,,9,,,,7,,
				4,,,,,,,9,
				8,,7,,,,,,
				,,,6,1,,,,
				,3,,,5,,,,1
				,1,,2,,,,7,5
				,,,1,4,3,,,`,
			want: &sudoku.Field{
				{7, 0, 0, 0, 0, 4, 8, 0, 0},
				{0, 0, 0, 0, 0, 5, 4, 0, 0},
				{0, 0, 9, 0, 0, 0, 7, 0, 0},
				{4, 0, 0, 0, 0, 0, 0, 9, 0},
				{8, 0, 7, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 6, 1, 0, 0, 0, 0},
				{0, 3, 0, 0, 5, 0, 0, 0, 1},
				{0, 1, 0, 2, 0, 0, 0, 7, 5},
				{0, 0, 0, 1, 4, 3, 0, 0, 0},
			},
			wantErr: false,
		},
		{
			name: "string in file sudoku file",
			input: `
				7,0,0, 0,0,4, 8,0,0
				0,0,0, 0,"test",5, 4,0,0
				0,0,9, 0,0,0, 7,0,0
				4,0,0, 0,0,0, 0,9,0
				8,0,7, 0,0,0, 0,0,0
				0,0,0, 6,1,0, 0,0,0
				0,3,0, 0,5,0, 0,0,1
				0,1,0, 2,0,0, 0,7,5
				0,0,0, 1,4,3, 0,0,0`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "to many lines",
			input: `
				7,0,0, 0,0,4, 8,0,0
				0,0,0, 0,0,5, 4,0,0
				0,0,9, 0,0,0, 7,0,0
				4,0,0, 0,0,0, 0,9,0
				8,0,7, 0,0,0, 0,0,0
				0,0,0, 6,1,0, 0,0,0
				0,3,0, 0,5,0, 0,0,1
				0,1,0, 2,0,0, 0,7,5
				0,1,0, 2,0,0, 0,7,5
				0,0,0, 1,4,3, 0,0,0`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "too little lines",
			input: `
				7,0,0, 0,0,4, 8,0,0
				0,0,0, 0,0,5, 4,0,0
				0,0,9, 0,0,0, 7,0,0
				4,0,0, 0,0,0, 0,9,0
				8,0,7, 0,0,0, 0,0,0
				0,0,0, 6,1,0, 0,0,0
				0,3,0, 0,5,0, 0,0,1
				0,0,0, 1,4,3, 0,0,0`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty file",
			input: "",
			want:    nil,
			wantErr: true,
		},
	}
	var r io.Reader
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r = strings.NewReader(tt.input)
			got, err := readCSV(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSV() = \n%v,\n want \n%v", got, tt.want)
			}
		})
	}
}
