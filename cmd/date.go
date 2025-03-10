package cmd

import "time"

func formatDate(date string) string {
	ref := "2006-01-02T15:04:05-07:00"
	t, err := time.Parse(ref, date)
	if err != nil {
		return ""
	}
	loc, err := time.LoadLocation(
		"Europe/Berlin")
	if err != nil {
		return ""
	}
	return t.In(loc).Format(
		"02.01.2006 um 15:04")
}
