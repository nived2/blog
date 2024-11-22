# Go Static Blog Generator

A lightweight, fast, and easy-to-use static blog generator written in Go. Perfect for hosting on GitHub Pages! Create your blog with markdown files and deploy it for free.

## ✨ Features

- 📝 Write posts in Markdown
- 🎨 Clean and responsive design using Water.css
- ⚡ Fast static site generation
- 🚀 One-command deployment to GitHub Pages
- 📱 Mobile-friendly layout
- 🔍 SEO-friendly output
- 💻 Cross-platform support

## 🚀 Quick Start

1. **Clone this repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/blog
   cd blog
   ```

2. **Install Go** (1.16 or later)
   - Download from [golang.org](https://golang.org/dl/)

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Add your blog posts**
   - Create `.md` files in the `content` directory
   - Use standard Markdown syntax

5. **Generate your site**
   ```bash
   go run main.go
   ```

6. **Deploy to GitHub Pages**
   - Push to your GitHub repository
   - Enable GitHub Pages in repository settings
   - Select the `public` directory as the source

## 📝 Writing Posts

1. Create a new `.md` file in the `content` directory
2. Write your post in Markdown format
3. Run the generator
4. Your post will be available at `your-site.com/post-name.html`

Example post:
```markdown
# My First Post

Welcome to my blog! This is a sample post.

## Features
- Easy to write
- Markdown support
- Code highlighting
```

## 🎨 Customization

### Templates
Edit the HTML templates in `templates/` directory:
- `index.html`: Blog homepage
- `post.html`: Individual post layout

### Styling
The default theme uses [Water.css](https://watercss.kognise.dev/) for a clean, responsive design.

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

Please read our [Contributing Guidelines](CONTRIBUTING.md) before submitting changes.

## 📜 License

This project is open source and available under the MIT License.

MIT License

Copyright (c) 2023 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## 🌟 Support

If you find this project helpful, please give it a star on GitHub!
