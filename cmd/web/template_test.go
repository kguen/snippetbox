package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{"UTC", time.Date(2021, 11, 01, 14, 0, 0, 0, time.UTC), "01 Nov 2021 at 14:00"},
		{"Empty", time.Time{}, ""},
		{"UTC+7", time.Date(2021, 11, 01, 14, 0, 0, 0, time.FixedZone("UTC+7", -7*60*60)), "01 Nov 2021 at 21:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}
