# Guidelines for Adding New Learn Articles

## Overview
This document provides step-by-step instructions for adding new articles to the Feral File Learn section. The Learn section is accessible at [feralfile.com/learn](http://feralfile.com/learn).

## Steps to Add a New Article

### 1. Create the Article File
1. Create a new Markdown file in the `learning/` directory
2. Use the following template structure at the beginning of your file:
   ```markdown
   ---
   title: "Your Article Title"
   description: "A brief description of your article"
   navTitle: "Short Title for Navigation"  # Optional - will use title if not specified
   weight: 1 # Optional - define the display order of your article under Learn section
   draft: "true" # Optional - default is false, your article will auto publish if not provided
   ---
   ```
3. Write your article content below the frontmatter
4. Commit directly to `main` branch (or open a PR).

### Example Article
For styling and structure reference, see the [How to Collect Digital Art](https://feralfile.com/learn/how-to-collect-digital-art) article.

## How It Works

The system automatically:
1. Detects changes to markdown files in the `learning/` directory
2. Parses the frontmatter from each file
3. Generates the `learn-manifest.json` file with the correct structure
4. Updates the website with your new content

## Important Notes
- The filename you choose will become the URL path for your article on the website
- Make sure to use descriptive, URL-friendly filenames
- The `title` and `description` in the frontmatter serve two purposes:
  1. SEO optimization for search engines
  2. Article preview in the list of all Learn pages (shown as a card with title and description)
- The `navTitle` in the frontmatter is used for the website header navigation (falls back to `title` if not specified)
- The order of articles in the manifest is determined by the `weight` property in each article's frontmatter
- After pushing your article, it will be automatically published to the Learn section of the website

## Design Decisions

### Why Automated Manifest Generation?
The decision to automate manifest generation was made for several reasons:

1. **Simplified Workflow**: 
   - Editors only need to focus on content creation
   - No need to manually update learn-manifest.json
   - Reduces the chance of errors in the manifest file

2. **Consistency**:
   - Ensures all articles follow the same structure
   - Automatically validates required frontmatter fields
   - Maintains consistent formatting across the site

3. **Maintainability**:
   - Single source of truth for article metadata
   - Easier to update the system without changing editor workflows
   - Centralized control over how articles are processed

## Best Practices
1. Use clear, descriptive titles in the frontmatter
2. Keep navTitle concise for navigation display
3. Keep descriptions concise but informative (they will be displayed in article previews)
4. Follow consistent formatting across all articles
5. Include relevant images and links where appropriate
6. Test all links before publishing
7. Review the article for typos and formatting issues
8. Use a consistent structure for all articles:
   - Start with an introduction
   - Use clear section headings
   - End with a conclusion or next steps
9. Include a table of contents for longer articles
10. Use callouts to highlight important information
11. Ensure images are high quality and properly sized
12. Use consistent terminology throughout the article
