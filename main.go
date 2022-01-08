package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	timestampBytes, err := os.ReadFile(".timestamp")
	if err != nil {
		timestampBytes = []byte{}
	}
	timestamp := string(timestampBytes[:len(timestampBytes)-1])

	htmlFiles, err := findFiles("templates")
	if err != nil {
		log.Fatal(err)
	}
	// create html files
	for _, tmpl := range htmlFiles {
		if tmpl == "_base.html" {
			continue
		}
		fmt.Printf("[INFO] Processing template %-20s", tmpl)
		t := template.Must(template.New("_base.html").ParseFiles(filepath.Join("templates", "_base.html"), filepath.Join("templates", tmpl)))
		f, err := os.Create(tmpl)
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(f, struct{ Timestamp string }{Timestamp: timestamp}); err != nil {
			log.Fatal(err)
		}
		fmt.Println("[DONE]")
	}

	images, err := findFiles(filepath.Join("static", "images"))
	if err != nil {
		log.Fatal(err)
	}
	darkimages, err := findFiles(filepath.Join("static", "images", "dark"))
	if err != nil {
		log.Fatal(err)
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

func findFiles(directory string) ([]string, error) {
	files := []string{}
	direntries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range direntries {
		if !d.IsDir() {
			i, err := d.Info()
			if err != nil {
				return nil, err
			}
			files = append(files, i.Name())
		}
	}
	return files, nil
}
