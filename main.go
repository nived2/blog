package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title       string
	Content     template.HTML
	Date        time.Time
	URLPath     string
	Description string
}

func main() {
	// Create necessary directories
	dirs := []string{"content", "templates", "public"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Generate the site
	if err := generateSite(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Site generated successfully!")
}

func generateSite() error {
	// Read and parse all markdown files
	posts, err := parseMarkdownFiles()
	if err != nil {
		return fmt.Errorf("failed to parse markdown files: %v", err)
	}

	// Generate HTML files
	if err := generateHTML(posts); err != nil {
		return fmt.Errorf("failed to generate HTML: %v", err)
	}

	return nil
}

func parseMarkdownFiles() ([]Post, error) {
	var posts []Post
	files, err := ioutil.ReadDir("content")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			content, err := ioutil.ReadFile(filepath.Join("content", file.Name()))
			if err != nil {
				return nil, err
			}

			// Parse markdown
			extensions := parser.CommonExtensions | parser.AutoHeadingIDs
			parser := parser.NewWithExtensions(extensions)
			html := markdown.ToHTML(content, parser, nil)

			// Create post
			post := Post{
				Title:   strings.TrimSuffix(file.Name(), ".md"),
				Content: template.HTML(html),
				Date:    time.Now(),
				URLPath: "/" + strings.TrimSuffix(file.Name(), ".md") + ".html",
			}
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func generateHTML(posts []Post) error {
	// Create templates directory if it doesn't exist
	if err := os.MkdirAll("public", 0755); err != nil {
		return err
	}

	// Generate individual post pages
	for _, post := range posts {
		if err := generatePostPage(post); err != nil {
			return err
		}
	}

	// Generate index page
	if err := generateIndexPage(posts); err != nil {
		return err
	}

	return nil
}

func generatePostPage(post Post) error {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join("public", strings.TrimPrefix(post.URLPath, "/")))
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, post)
}

func generateIndexPage(posts []Post) error {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return err
	}

	file, err := os.Create("public/index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, posts)
}
