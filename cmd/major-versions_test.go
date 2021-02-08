package cmd

import (
	"testing"
)

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
		{
			name:       "EA",
			args:       argsStruct{"ea"},
			wantResult: "latest_ea",
			wantErr:    false,
		},
		{
			name:       "STS",
			args:       argsStruct{"sts"},
			wantResult: "latest_sts",
			wantErr:    false,
		},
		{
			name:       "MTS",
			args:       argsStruct{"mts"},
			wantResult: "latest_mts",
			wantErr:    false,
		},
		{
			name:       "LTS",
			args:       argsStruct{"lts"},
			wantResult: "latest_lts",
			wantErr:    false,
		},
		{
			name:       "INVALID",
			args:       argsStruct{"alts"},
			wantResult: "",
			wantErr:    true,
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
