package sudoku

import (
	"reflect"
	"testing"
)

func TestSolverField_String(t *testing.T) {
	tests := []struct {
		name string
		c    SolverField
		want string
	}{
		{
			name: "empty possibilities",
			c:    SolverField{},
			want: "",
		},
		{
			name: "with empty possibilities",
			c:    SolverField{{&Possibilities{}}},
			want: "(0,0): []",
		},
		{
			name: "with possibilities",
			c: SolverField{
				{
					&Possibilities{true, false, true, false, false},
					&Possibilities{false, false, true, false, true},
				},
			},
			want: "(0,0): [1,3]\n(1,0): [3,5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("SolverField.String() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

var testField = &Field{
	{7, 0, 0, 0, 0, 4, 8, 0, 0},
	{0, 0, 0, 0, 0, 5, 4, 0, 0},
	{0, 0, 9, 0, 0, 0, 7, 0, 0},

	{4, 0, 0, 0, 0, 0, 0, 9, 0},
	{8, 0, 7, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 6, 1, 0, 0, 0, 0},

	{0, 3, 0, 0, 5, 0, 0, 0, 1},
	{0, 1, 0, 2, 0, 0, 0, 7, 5},
	{0, 0, 0, 1, 4, 3, 0, 0, 0},
}

func TestField_PossibilitiesInRow(t *testing.T) {
	tests := []struct {
		name string
		f    *Field
		row  int
		want *Possibilities
	}{
		{
			name: "empty field",
			f:    &Field{},
			row:  0,
			want: NewPossibilities(),
		}, {
			name: "row with numbers field",
			f:    testField,
			row:  0,
			want: &Possibilities{true, true, true, false, true, true, false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.PossibilitiesInRow(tt.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.PossibilitiesInRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_PossibilitiesInColumn(t *testing.T) {
	tests := []struct {
		name string
		f    *Field
		col  int
		want *Possibilities
	}{
		{
			name: "empty field",
			f:    &Field{},
			col:  0,
			want: NewPossibilities(),
		},
		{
			name: "row with numbers field",
			f:    testField,
			col:  3,
			want: &Possibilities{false, false, true, true, true, false, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.PossibilitiesInColumn(tt.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.PossibilitiesInColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_PossibilitiesInSquare(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		f    *Field
		args args
		want *Possibilities
	}{
		{
			name: "empty field",
			f:    &Field{},
			args: args{},
			want: NewPossibilities(),
		},
		{
			name: "test field",
			f:    testField,
			args: args{x: 3, y: 6},
			want: &Possibilities{false, false, false, false, false, true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.PossibilitiesInSquare(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.PossibilitiesInSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

var testField2 = &Field{
	{7, 0, 0, 0, 0, 4, 8, 0, 0},
	{0, 0, 0, 0, 0, 5, 4, 0, 0},
	{0, 0, 9, 0, 0, 0, 7, 0, 0},

	{4, 0, 0, 0, 0, 0, 0, 9, 0},
	{8, 0, 7, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 3, 0, 0, 0, 0, 0, 0, 1},
	{0, 1, 0, 2, 0, 0, 0, 7, 5},
	{0, 0, 0, 1, 0, 3, 0, 0, 0},
}
var testFieldSolved = &Field{
	{7, 2, 5, 3, 6, 4, 8, 1, 9},
	{1, 8, 3, 9, 7, 5, 4, 6, 2},
	{6, 4, 9, 8, 2, 1, 7, 5, 3},

	{4, 6, 1, 5, 3, 8, 2, 9, 7},
	{8, 5, 7, 4, 9, 2, 1, 3, 6},
	{3, 9, 2, 6, 1, 7, 5, 8, 4},

	{2, 3, 8, 7, 5, 9, 6, 4, 1},
	{9, 1, 4, 2, 8, 6, 3, 7, 5},
	{5, 7, 6, 1, 4, 3, 9, 2, 8},
}

func TestSolve(t *testing.T) {
	type args struct {
		f        Field
		onUpdate UpdateFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *Field
		wantErr bool
	}{
		{
			name: "empty field",
			args: args{
				f: Field{},
			},
			want:    &Field{},
			wantErr: true,
		},
		{
			name: "solve field",
			args: args{
				f: *testField,
			},
			want:    testFieldSolved,
			wantErr: false,
		},
		{
			name: "no clear solution ",
			args: args{
				f: *testField2,
			},
			want: &Field{
				{7, 0, 0, 0, 0, 4, 8, 0, 0},
				{0, 8, 0, 0, 0, 5, 4, 0, 0},
				{0, 4, 9, 0, 0, 0, 7, 0, 0},

				{4, 0, 0, 0, 0, 0, 0, 9, 0},
				{8, 0, 7, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},

				{0, 3, 0, 0, 0, 0, 0, 0, 1},
				{0, 1, 0, 2, 0, 0, 3, 7, 5},
				{0, 7, 0, 1, 0, 3, 0, 0, 0},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.args.f, tt.args.onUpdate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
