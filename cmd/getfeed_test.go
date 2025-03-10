package cmd

import "testing"

func Test_getId(t *testing.T) {
	type args struct {
		guid string
	}
	tests := []struct {
		name    string
		guid    string
		want    string
		wantErr bool
	}{
		{"ok", "http://heise.de/-1234", "1234", false},
		{"no number", "http://heise.de/-", "", false},
		{"empty", "", "", true},
		{"other url", "https://heise.de", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getId(tt.guid)
			if tt.wantErr != (err != nil) {
				t.Errorf("getId() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("getId() = %v, want %v", got, tt.want)
			}
		})
	}
}
