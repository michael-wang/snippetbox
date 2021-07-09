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
		{
			name: "UTC",
			tm:   time.Date(2021, 7, 9, 6, 16, 0, 0, time.UTC),
			want: "09 Jul 2021 at 06:16",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "TPE",
			tm:   time.Date(2021, 7, 9, 14, 16, 0, 0, time.FixedZone("TPE", 8*60*60)),
			want: "09 Jul 2021 at 06:16",
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			hd := humanDate(c.tm)
			if hd != c.want {
				t.Errorf("want %q; got: %q", c.want, hd)
			}
		})
	}
}
