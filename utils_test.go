package goldtest

import (
	"testing"
)

func Test_removeSpecialCharacters(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		path string
		want string
	}{
		{"Clean string", "test", "test"},
		{"Clean string - 1", "TeSt", "TeSt"},
		{"Underscore", "test_with_underscore", "test_with_underscore"},
		{"Dash", "test-with-dash", "test-with-dash"},
		{"Special characters", "test!@#$%^", "test"},
		{"Spaces", "test with spaces", "test_with_spaces"},
		{"Slashes", "test/with/slash", "test/with/slash"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSpecialCharacters(tt.path); got != tt.want {
				t.Errorf("removeSpecialCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeRead(t *testing.T) {
	type args struct {
		fileName string
		bytes    []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Write file to directory", args{"testdir/testfile_tmp", []byte("test")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeFile(tt.args.fileName, tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("writeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if b, _ := readFile(tt.args.fileName); string(b) != string(tt.args.bytes) {
				t.Errorf("readFile() get wrong output = %v, want %v", string(b), string(tt.args.bytes))
			}
		})
	}
}
