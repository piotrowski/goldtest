package goldtest

import (
	"testing"
)

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
