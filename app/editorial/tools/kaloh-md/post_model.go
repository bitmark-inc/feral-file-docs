// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2023 Bitmark Inc.

package main

import "time"

type KalohPost struct {
	ID                 int       `json:"id"`
	PublicationID      int       `json:"publication_id"`
	Type               string    `json:"type"`
	Title              string    `json:"title"`
	SocialTitle        string    `json:"social_title"`
	SectionID          int       `json:"section_id"`
	Subtitle           string    `json:"subtitle"`
	Slug               string    `json:"slug"`
	PostDate           time.Time `json:"post_date"`
	Audience           string    `json:"audience"`
	CoverImage         string    `json:"cover_image"`
	CoverImageIsSquare bool      `json:"cover_image_is_square"`
	Description        string    `json:"description"`
	BodyHTML           string    `json:"body_html"`
	TruncatedBodyText  string    `json:"truncated_body_text"`
	Wordcount          int       `json:"wordcount"`
	CommentCount       int       `json:"comment_count"`
}
