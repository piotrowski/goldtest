package goldtest

import (
	"testing"
)

func TestAssert(t *testing.T) {
	type testStruct struct {
		UUID    string
		name    string
		value   int64
		numbers []int8
	}

	type args struct {
		t        *testing.T
		actual   interface{}
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Bytes", args{t, []byte("test"), "testdir/testfile_byte"}},
		{"Number", args{t, 123456789, "testdir/testfile_number"}},
		{"Array", args{t, [4]int{1, 2, 3, 4}, "testdir/testfile_array"}},
		{"map", args{t, map[string]string{"test": "testing"}, "testdir/testfile_map"}},
		{"struct", args{t,
			testStruct{
				"123e4567-e89b-12d3-a456-426652340000",
				"apiotrowski312",
				9223372036854775807,
				[]int8{100, 120, 125, 10},
			}, "testdir/testfile_struct"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Assert(tt.args.t, tt.args.actual, tt.args.fileName)
		})
	}
}

func TestAssertJSON(t *testing.T) {
	type testStruct struct {
		UUID    string
		name    string
		value   int64
		numbers []int8
	}

	type args struct {
		t        *testing.T
		actual   interface{}
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Bytes", args{t, []byte("test"), "testdir/json_testfile_byte"}},
		{"Number", args{t, 123456789, "testdir/json_testfile_number"}},
		{"Array", args{t, [4]int{1, 2, 3, 4}, "testdir/json_testfile_array"}},
		{"map", args{t, map[string]string{"test": "testing"}, "testdir/json_testfile_map"}},
		{"struct", args{t,
			testStruct{
				"123e4567-e89b-12d3-a456-426652340000",
				"apiotrowski312",
				9223372036854775807,
				[]int8{100, 120, 125, 10},
			}, "testdir/json_testfile_struct"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertJSON(tt.args.t, tt.args.actual, tt.args.fileName)
		})
	}
}
