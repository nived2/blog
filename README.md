# Static Blog Generator

A lightweight static blog generator written in Go, perfect for hosting on GitHub Pages.

## Features

- Markdown support
- Clean and responsive design using Water.css
- Static site generation
- Easy to deploy to GitHub Pages

## Getting Started

1. Install Go (1.16 or later)
2. Clone this repository
3. Run `go mod tidy` to install dependencies

## Usage

1. Add your Markdown posts to the `content` directory
2. Run the generator:
   ```bash
   go run main.go
   ```
3. The generated site will be in the `public` directory

## Writing Posts

Create Markdown files in the `content` directory. The filename will be used as the URL slug.

Example: `content/hello-world.md` will become `your-site.com/hello-world.html`

## Deployment to GitHub Pages

1. Push your repository to GitHub
2. Enable GitHub Pages in your repository settings
3. Set the GitHub Pages source to the `public` directory
4. Your blog will be available at `your-username.github.io/repository-name`

## Customization

- Edit the templates in the `templates` directory to customize the HTML structure
- Modify the CSS in the template files to change the appearance

## License

MIT License
