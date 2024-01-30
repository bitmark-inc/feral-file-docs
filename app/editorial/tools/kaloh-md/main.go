// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2023 Bitmark Inc.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/microcosm-cc/bluemonday"
)

var (
	sanitizer = bluemonday.UGCPolicy()
	converter = md.NewConverter("", true, nil)
)

func getMardown(ctx context.Context, url string) (string, error) {

	apiURL := strings.ReplaceAll(url, "/p/", "/api/v1/posts/")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var post KalohPost
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		return "", err
	}

	html := sanitizer.Sanitize(post.BodyHTML)

	markdown, err := converter.ConvertString(html)
	if err != nil {
		return "", err
	}

	markdown = fmt.Sprintf("# %s\n\n%s\n\n---\n\n%s", post.Title, post.Description, markdown)

	return markdown, nil
}

func main() {
	sanitizer.AllowImages()
	sanitizer.AllowLists()
	sanitizer.AllowStandardURLs()
	sanitizer.AllowTables()

	if len(os.Args)-1 != 2 {
		fmt.Printf("Invalid args. Len: %s, Expected: 2. kaloh url output_file", os.Args)
		return
	}

	url := os.Args[1]
	filename := os.Args[2]

	ctx := context.Background()
	md, err := getMardown(ctx, url)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(md)
	if err != nil {
		panic(err)
	}
}
