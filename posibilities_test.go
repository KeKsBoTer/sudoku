package sudoku

import (
	"reflect"
	"testing"
)

func TestNewPossibilities(t *testing.T) {
	tests := []struct {
		name string
		want *Possibilities
	}{
		{
			name: "all numbers possible",
			want: &Possibilities{true, true, true, true, true, true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPossibilities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPossibilities_IsPossible(t *testing.T) {
	tests := []struct {
		name string
		p    *Possibilities
		n    int
		want bool
	}{
		{
			name: "number is possible",
			p:    &Possibilities{false, true},
			n:    2,
			want: true,
		}, {
			name: "number is not possible",
			p:    &Possibilities{false, false, false},
			n:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsPossible(tt.n); got != tt.want {
				t.Errorf("Possibilities.IsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPossibilities_OnlyOne(t *testing.T) {
	tests := []struct {
		name  string
		p     *Possibilities
		want  bool
		want1 int
	}{
		{
			name:  "has only one",
			p:     &Possibilities{false, false, false, true},
			want:  true,
			want1: 4,
		},
		{
			name:  "has more than one",
			p:     &Possibilities{false, true, false, true},
			want:  false,
			want1: 0,
		},
		{
			name:  "has no possiblities",
			p:     &Possibilities{},
			want:  false,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.OnlyOne()
			if got != tt.want {
				t.Errorf("Possibilities.OnlyOne() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Possibilities.OnlyOne() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPossibilities_Remove(t *testing.T) {
	tests := []struct {
		name string
		p    *Possibilities
		n    int
	}{
		{
			name: "remove possible number",
			p:    &Possibilities{false, true},
			n:    2,
		},
		{
			name: "remove not possible number",
			p:    &Possibilities{},
			n:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Remove(tt.n)
			if tt.p.IsPossible(tt.n) {
				t.Errorf("Possibilities.IsPossible() = true, want false")
			}
		})
	}
}

func TestPossibilities_Add(t *testing.T) {
	tests := []struct {
		name string
		p    *Possibilities
		n    int
	}{
		{
			name: "add possible number",
			p:    &Possibilities{false, true},
			n:    2,
		},
		{
			name: "add not possible number",
			p:    &Possibilities{},
			n:    3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Add(tt.n)
			if !tt.p.IsPossible(tt.n) {
				t.Errorf("Possibilities.IsPossible() = false, want true")
			}
		})
	}
}

func TestPossibilities_Empty(t *testing.T) {
	tests := []struct {
		name string
		p    *Possibilities
		want bool
	}{
		{
			name: "empty",
			p:    &Possibilities{},
			want: true,
		},
		{
			name: "not empty",
			p:    &Possibilities{false, true},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Empty(); got != tt.want {
				t.Errorf("Possibilities.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPossibilities_String(t *testing.T) {
	tests := []struct {
		name string
		p    Possibilities
		want string
	}{
		{
			name: "empty",
			p:    Possibilities{},
			want: "[]",
		},
		{
			name: "not empty",
			p:    Possibilities{true, false, false, true},
			want: "[1,4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Possibilities.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
