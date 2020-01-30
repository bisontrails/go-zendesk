package zendesk

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type Article struct {
	ID                *int64     `json:"id,omitempty"`
	URL               *string    `json:"url,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	Title             *string    `json:"title,omitempty"`
	Body              *string    `json:"body,omitempty"`
	Locale            *string    `json:"locale,omitempty"`
	SourceLocale      *string    `json:"source_locale,omitempty"`
	AuthorID          *int64     `json:"author_id,omitempty"`
	CommentsDisabled  *bool      `json:"comments_disabled,omitempty"`
	OutdatedLocales   *[]string  `json:"outdated_locales,omitempty"`
	Outdated          *bool      `json:"outdated,omitempty"`
	LabelNames        *[]string  `json:"label_names,omitempty"`
	Draft             *bool      `json:"draft,omitempty"`
	Promoted          *bool      `json:"promoted,omitempty"`
	Position          *int64     `json:"position,omitempty"`
	VoteSum           *int64     `json:"vote_sum,omitempty"`
	VoteCount         *int64     `json:"vote_count,omitempty"`
	SectionID         *int64     `json:"section_id,omitempty"`
	UserSegmentID     *int64     `json:"user_segment_id,omitempty"`
	PermissionGroupID *int64     `json:"permission_group_id,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	EditedAt          *time.Time `json:"edited_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

// ListArticlesOptions specifies the optional parameters for the list users methods.
type ListArticlesOptions struct {
	ListOptions

	StartTime  *time.Time `url:"start_time,omitempty"`
	LabelNames []string   `url:"label_names,comma,omitempty"`
}

// ListArticles fetches articles.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/attachments#getting-attachments
func (c *client) ListArticles(locale string, opts *ListArticlesOptions) ([]Article, error) {
	params, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = c.get(fmt.Sprintf("/api/v2/help_center/%s/articles.json?%s", locale, params.Encode()), out)
	return out.Articles, err
}
