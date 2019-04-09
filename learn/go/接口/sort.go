package main

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
	"sort"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var tracks = []*Track{
	{"go", "deliah", "from the root up", 2012, length("3m38s")},
	{"go1", "deliah1", "from the root up1", 2012, length("3m38s")},
	{"go ahead", "deliah kes", "from the root up df", 2007, length("8m")},
}

func length(v string) time.Duration {
	d, _ := time.ParseDuration(v)
	return d
}

func printTracks(ts []*Track)  {
	//const format = "%-10v\t%-10v\t%-20v\t%-10v\t%-10v\t\n"
	//fmt.Printf(format, "title", "artist", "aluble", "year", "length")
	//fmt.Printf(format, "-----", "------", "------", "----", "------")
	//
	//for _, t := range ts{
	//	fmt.Printf(format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	//}

	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "title", "artist", "aluble", "year", "length")
	fmt.Fprintf(tw, format, "-----", "------", "------", "----", "------")
	for _, t := range ts {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 计算各列宽度并输出表格
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int)  {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	t []*Track
	less func(x, y *Track) bool
}

func (x customSort)Len() int {
	return len(x.t)
}

func (x customSort)Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort)Swap(i, j int)  {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func main()  {
	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	// 因为这里的data实现了 Sort 的所有接口
	sort.Sort(customSort{tracks, func(x, y *Track) bool {

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

	printTracks(tracks)


	value := []int{4,21,6,7,2,8,1,8,0}
	// 顺序排序
	sort.Ints(value)
	fmt.Println(value)
	//逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(value)))  // sort.IntSlice() 强制转换
	fmt.Println(value)
	// 判断是否排序过
	fmt.Println(sort.IntsAreSorted(value))
}
