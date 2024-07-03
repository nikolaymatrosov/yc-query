package filter

import "testing"

func TestIn(t *testing.T) {
	type args struct {
		field  string
		values []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				field:  "name",
				values: []string{},
			},
			want: "",
		},
		{
			name: "single",
			args: args{
				field:  "name",
				values: []string{"value"},
			},
			want: "name IN ('value')",
		},
		{
			name: "multiple",
			args: args{
				field:  "name",
				values: []string{"value1", "value2"},
			},
			want: "name IN ('value1', 'value2')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.field, tt.args.values...); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}
