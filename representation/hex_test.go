package representation

import "testing"

func TestHexConvert(t *testing.T) {
	type args struct {
		decimal uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "115 decimal",
			args: args{decimal: 115},
			want: "73",
		},
		{
			name: "255 decimal",
			args: args{decimal: 255},
			want: "FF",
		},
		{
			name: "0 decimal",
			args: args{decimal: 0},
			want: "00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hex(tt.args.decimal); got != tt.want {
				t.Errorf("HexConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
