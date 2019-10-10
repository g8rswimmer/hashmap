package hashmap

import "testing"

func Test_defaultEqualer_Equals(t *testing.T) {
	type args struct {
		obj1 interface{}
		obj2 interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Equals",
			args: args{
				obj1: "they are equal",
				obj2: "they are equal",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Not Equals",
			args: args{
				obj1: 10,
				obj2: 12,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultEqualer{}
			got, err := d.Equals(tt.args.obj1, tt.args.obj2)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultEqualer.Equals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("defaultEqualer.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
