package cmd

import "testing"

func Test_fromShortToLatest(t *testing.T) {
	type argsStruct struct {
		value string
	}
	tests := []struct {
		name       string
		args       argsStruct
		wantResult string
		wantErr    bool
	}{
		{
			name:       "GA",
			args:       argsStruct{"ga"},
			wantResult: "latest_ga",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := fromShortToLatest(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromShortToLatest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("fromShortToLatest() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
