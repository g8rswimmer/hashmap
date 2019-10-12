package hashmap

import (
	"reflect"
	"testing"
)

func Test_HashMap_Put(t *testing.T) {
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
			h := &HashMap{
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

func TestHashMap_Get(t *testing.T) {
	type fields struct {
		hash  Hasher
		equal Equaler
		table []entries
		size  uint64
	}
	type args struct {
		k interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "One Entry",
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
			},
			want:    "yep",
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
						entry{
							key: uint64(32),
							obj: "yeppers",
						},
						entry{
							key: uint64(42),
							obj: "yep!!!",
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
				k: uint64(32),
			},
			want:    "yeppers",
			wantErr: false,
		},
		{
			name: "Can't find Entry",
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
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashMap{
				hash:  tt.fields.hash,
				equal: tt.fields.equal,
				table: tt.fields.table,
				size:  tt.fields.size,
			}
			got, err := h.Get(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashMap.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashMap(t *testing.T) {
	type args struct {
		size  uint64
		hash  Hasher
		equal Equaler
	}
	tests := []struct {
		name string
		args args
		want *HashMap
	}{
		{
			name: "New",
			args: args{
				size:  10,
				hash:  &mockHasher{},
				equal: &mockEqualer{},
			},
			want: &HashMap{
				hash:  &mockHasher{},
				equal: &mockEqualer{},
				table: make([]entries, 10),
				size:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMap(tt.args.size, tt.args.hash, tt.args.equal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
