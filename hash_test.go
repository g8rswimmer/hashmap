package hashmap

import (
	"reflect"
	"testing"
)

func Test_hashMap_Put(t *testing.T) {
	type fields struct {
		hash  Hasher
		equal Equaler
		table []entries
		size  uint64
	}
	type args struct {
		k interface{}
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entries
		wantErr bool
	}{
		{
			name: "First Entry",
			fields: fields{
				hash:  &mockHasher{},
				equal: &mockEqualer{},
				table: make([]entries, 10),
				size:  10,
			},
			args: args{
				k: uint64(12),
				v: "yep",
			},
			want: []entries{
				nil,
				nil,
				entries{
					entry{
						key: uint64(12),
						obj: "yep",
					},
				},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
			wantErr: false,
		},
		{
			name: "Hash Collision Entry",
			fields: fields{
				hash:  &mockHasher{},
				equal: &mockEqualer{},
				table: []entries{
					nil,
					nil,
					entries{
						entry{
							key: uint64(12),
							obj: "yep",
						},
					},
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
				},
				size: 10,
			},
			args: args{
				k: uint64(22),
				v: "yeppers",
			},
			want: []entries{
				nil,
				nil,
				entries{
					entry{
						key: uint64(12),
						obj: "yep",
					},
					entry{
						key: uint64(22),
						obj: "yeppers",
					},
				},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
			wantErr: false,
		},
		{
			name: "Replace Entry",
			fields: fields{
				hash:  &mockHasher{},
				equal: &mockEqualer{},
				table: []entries{
					nil,
					nil,
					entries{
						entry{
							key: uint64(12),
							obj: "yep",
						},
					},
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
				},
				size: 10,
			},
			args: args{
				k: uint64(12),
				v: "yeppers",
			},
			want: []entries{
				nil,
				nil,
				entries{
					entry{
						key: uint64(12),
						obj: "yeppers",
					},
				},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
			wantErr: false,
		},
		{
			name: "Multiple Entry",
			fields: fields{
				hash:  &mockHasher{},
				equal: &mockEqualer{},
				table: []entries{
					nil,
					nil,
					entries{
						entry{
							key: uint64(12),
							obj: "yep",
						},
					},
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
				},
				size: 10,
			},
			args: args{
				k: uint64(33),
				v: "yeppers",
			},
			want: []entries{
				nil,
				nil,
				entries{
					entry{
						key: uint64(12),
						obj: "yep",
					},
				},
				entries{
					entry{
						key: uint64(33),
						obj: "yeppers",
					},
				},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			},
			wantErr: false,
		},
		{
			name: "Hash Error",
			fields: fields{
				hash:  &mockHasher{},
				equal: &mockEqualer{},
				table: make([]entries, 10),
				size:  10,
			},
			args: args{
				k: "uint",
				v: "yep",
			},
			want:    make([]entries, 10),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashMap{
				hash:  tt.fields.hash,
				equal: tt.fields.equal,
				table: tt.fields.table,
				size:  tt.fields.size,
			}
			if err := h.Put(tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("hashMap.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(h.table, tt.want) == false {
				t.Errorf("hashMap.Put() got = %v, want %v", h.table, tt.want)
			}
		})
	}
}
