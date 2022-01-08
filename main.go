package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
)

var (
	tmpls = []string{"index.html", "progress.html", "demo.html"}
)

func main() {
	for _, tmpl := range tmpls {
		t := template.Must(template.New("_base.html").ParseFiles(filepath.Join("templates", "_base.html"), filepath.Join("templates", tmpl)))
		f, err := os.Create(tmpl)
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(f, nil)
	}

	dir, err := os.ReadDir(filepath.Join("static", "images"))
	if err != nil {
		log.Fatal(err)
	}
	images := []string{}
	darkimages := []string{}
	for _, d := range dir {
		// fmt.Println(dir)
		if !d.IsDir() {
			i, err := d.Info()
			if err != nil {
				log.Fatal(err)
			}
			images = append(images, i.Name())
		}
	}
	darkdir, err := os.ReadDir(filepath.Join("static", "images", "dark"))
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range darkdir {
		// fmt.Println(dir)
		if !d.IsDir() {
			i, err := d.Info()
			if err != nil {
				log.Fatal(err)
			}
			darkimages = append(darkimages, i.Name())
		}
	}
	sort.Strings(images)
	// sort.Sort(sort.Reverse(sort.StringSlice(darkimages)))
	for _, darkimage := range darkimages {
		i := sort.SearchStrings(images, darkimage)
		if images[i] == darkimage && i < len(images) {
			images = append(images[:i], images[i+1:]...)
		}
	}

	for _, image := range images {
		fmt.Println("[WARN] no dark variant found for", image)
	}

}
