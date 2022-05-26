package hello_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/osekihan-7s/go-hello"
)

type a struct {
	name string
	age  int
}

func TestShello(t *testing.T) {

	type arg struct {
		name string
	}

	tests := []struct {
		name string
		arg  arg
		want string
	}{
		{
			name: "Alice",
			arg:  arg{name: "Alice"},
			want: "Hello, Alice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hello.Shello(tt.name)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("return value mismatch (-want, +got)\n%v", diff)
			}

		})
	}

}
