/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/
package main

import "time"
import "fmt"

func main() {
	// cheating, use time package
	week := 7 * 24 * time.Hour
	end := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)
	count := 0
	// start is 1901-01-06, first Sunday
	for start := time.Date(1901, time.January, 6, 0, 0, 0, 0, time.UTC); start.Before(end) ; start = start.Add(week) {
		if start.Day() == 1 {
			count++
		}
	}
	fmt.Println(count)
}