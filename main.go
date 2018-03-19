package main

import (
	"learn/ch8/thumbnail"
	"fmt"
	"runtime"
)

func main()  {
	const basePath = "/home/kitralzhong/tmp/gangan"
	files := []string{
		"gangan12.jpg",
		"gangan1.jpg",
	}
	for i:=1; i<10;i++{
		err := MakeThumbnails4(files, basePath)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("%d %d\n", i, runtime.NumGoroutine())
	}
}



func MakeThumbnails4(files []string, base string) error  {
	// ? 书上说这里能够让goroutine正常退出 但是没效果
	ch := make(chan error, len(files))
	for _, f := range files{
		go func(f string) {
			_, err := thumbnail.ImageFile(base + "/" + f)
			ch <- err
		}(f)
	}
	for range files{
		if err := <- ch; err != nil {
			return err
		}
	}
	return nil
}
