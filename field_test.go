package sudoku

import (
	"testing"
)

func TestField_EmptyCells(t *testing.T) {
	tests := []struct {
		name string
		f    *Field
		want int
	}{
		{
			name: "empty",
			f:    &Field{},
			want: 9 * 9,
		}, {
			name: "not empty",
			f:    testField,
			want: 9*9 - 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.EmptyCells(); got != tt.want {
				t.Errorf("Field.EmptyCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_Check(t *testing.T) {
	// generate field with error
	errorfield := *testField
	errorfield[0][1] = 7

	tests := []struct {
		name    string
		f       *Field
		wantErr bool
	}{
		{
			name:    "with zeros",
			f:       testField,
			wantErr: false,
		},
		{
			name:    "solved field",
			f:       testFieldSolved,
			wantErr: false,
		},
		{
			name:    "field with errors",
			f:       &errorfield,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Check(); (err != nil) != tt.wantErr {
				t.Errorf("Field.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
