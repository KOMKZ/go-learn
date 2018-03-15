package main

import (
	"learn/ch7/sorting"
	"time"
	"sort"
	"fmt"
)

func length(s string) time.Duration  {
	d, err := time.ParseDuration(s)
	if err != nil{
		panic(s)
	}
	return d
}



func main()  {
	tracks := []*sorting.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	// normal output
	sorting.PrintTracks(tracks)
	// sort by artist
	fmt.Println()
	sort.Sort(sorting.ByArtist(tracks))
	sorting.PrintTracks(tracks)
	sort.Sort(sorting.CustomSort{tracks, func(x, y *sorting.Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	fmt.Println()
	sorting.PrintTracks(tracks)

}


