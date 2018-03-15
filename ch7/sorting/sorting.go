package sorting

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type ByArtist []*Track

func (a ByArtist)Len() int {
	return len(a)
}

func (a ByArtist)Less(i, j int) bool  {
	return a[i].Artist < a[j].Artist
}

func (a ByArtist)Swap(i, j int) {
	a[i].Artist, a[j].Artist = a[j].Artist, a[i].Artist
}

type CustomSort struct {
	T []*Track
	CustomLess func(x, y *Track) bool
}

func (c CustomSort)Len() int  {
	return len(c.T)
}
func (c CustomSort)Less(i, j int) bool  {
	return c.CustomLess(c.T[i], c.T[j])
}
func (c CustomSort)Swap(i, j int) {
	c.T[i], c.T[j] = c.T[j], c.T[i]
}



func PrintTracks(tracks []*Track)  {
	format := "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format,"Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _,t := range tracks{
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
