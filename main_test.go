package main

import (
	"testing"
)

func Test_slugify(t *testing.T) {
	channel := make(chan string)
	type args struct {
		stringToSlugify string
		c chan string
	}
	tests := []struct {
		name string
		args args
		expected string
	}{
		{"empty slug",args{stringToSlugify: "",c: channel},""},
		{"lowercase",args{stringToSlugify: "TEST",c: channel},"test"},
		{"ascii",args{stringToSlugify: "ääaa",c: channel},"aeaeaa"},
		{"space",args{stringToSlugify: "    ",c: channel},"____"},
		{"dots",args{stringToSlugify: "Áäää  .asd.T",c: channel},"aaeaeae__.asd.t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slugify(tt.args.stringToSlugify, tt.args.c)
		})
		result := <- tt.args.c
		if tt.expected != result {
			t.Errorf("slugify error on case %s, got %v, expected %v",tt.name,result,tt.expected)
		}
	}
}
