package helper

import "testing"

func Test_fromBoolToYesNo(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Boolean true",
			args: args{true},
			want: "Yes",
		},
		{
			name: "Boolean false",
			args: args{false},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBoolToYesNo(tt.args.value); got != tt.want {
				t.Errorf("fromBoolToYesNo() = %v, want %v", got, tt.want)
			}
		})
	}
}
