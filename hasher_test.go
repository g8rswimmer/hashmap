package hashmap

import (
	"testing"
)

func Test_defaultHasher_intHash(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "int positive",
			args: args{
				obj: int(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "int negative",
			args: args{
				obj: int(-10),
			},
			want:    18446744073709551606,
			wantErr: false,
		},
		{
			name: "int8 positive",
			args: args{
				obj: int8(10),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "int8 negative",
			args: args{
				obj: int8(-10),
			},
			want:    18446744073709551606,
			wantErr: false,
		},
		{
			name: "int16 positive",
			args: args{
				obj: int16(10000),
			},
			want:    10000,
			wantErr: false,
		},
		{
			name: "int16 negative",
			args: args{
				obj: int16(-10000),
			},
			want:    18446744073709541616,
			wantErr: false,
		},
		{
			name: "int32 positive",
			args: args{
				obj: int32(10000000),
			},
			want:    10000000,
			wantErr: false,
		},
		{
			name: "int32 negative",
			args: args{
				obj: int32(-10000000),
			},
			want:    18446744073699551616,
			wantErr: false,
		},
		{
			name: "int64 positive",
			args: args{
				obj: int64(8589934592),
			},
			want:    8589934592,
			wantErr: false,
		},
		{
			name: "int64 negative",
			args: args{
				obj: int64(-8589934592),
			},
			want:    18446744065119617024,
			wantErr: false,
		},
		{
			name: "uint",
			args: args{
				obj: uint(100),
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "uint8",
			args: args{
				obj: uint8(100),
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "uint16",
			args: args{
				obj: uint16(20000),
			},
			want:    20000,
			wantErr: false,
		},
		{
			name: "uint32",
			args: args{
				obj: uint32(20000000),
			},
			want:    20000000,
			wantErr: false,
		},
		{
			name: "uint64",
			args: args{
				obj: uint64(9589934592),
			},
			want:    9589934592,
			wantErr: false,
		},
		{
			name: "Error",
			args: args{
				obj: "error",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultHasher{}
			got, err := d.intHash(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultHasher.intHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("defaultHasher.intHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultHasher_stringHash(t *testing.T) {
	type args struct {
		obj string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "one letter string",
			args: args{
				obj: "a",
			},
			want:    97,
			wantErr: false,
		},
		{
			name: "two letter string",
			args: args{
				obj: "ab",
			},
			want:    293,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultHasher{}
			got, err := d.stringHash(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultHasher.stringHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("defaultHasher.stringHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultHasher_floatHash(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "float32 positive",
			args: args{
				obj: float32(100.4343),
			},
			want:    1120460381,
			wantErr: false,
		},
		{
			name: "float32 negative",
			args: args{
				obj: float32(-100.4343),
			},
			want:    3267944029,
			wantErr: false,
		},
		{
			name: "float64 positive",
			args: args{
				obj: float64(8589934592.338),
			},
			want:    4755801206503420985,
			wantErr: false,
		},
		{
			name: "float64 negative",
			args: args{
				obj: float64(-8589934592.4343),
			},
			want:    13979173243358247282,
			wantErr: false,
		},
		{
			name: "Error",
			args: args{
				obj: "error",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultHasher{}
			got, err := d.floatHash(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultHasher.floatHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("defaultHasher.floatHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultHasher_Hash(t *testing.T) {
	type s struct {
		One int
		Str string
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "int",
			args: args{
				obj: int(100),
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "two letter string",
			args: args{
				obj: "ab",
			},
			want:    293,
			wantErr: false,
		},
		{
			name: "float32 negative",
			args: args{
				obj: float32(-100.4343),
			},
			want:    3267944029,
			wantErr: false,
		},
		{
			name: "bool",
			args: args{
				obj: true,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Pass",
			args: args{
				obj: s{
					One: 12,
					Str: "something",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultHasher{}
			got, err := d.Hash(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultHasher.Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("defaultHasher.Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
