package yonoma

import (
	"fmt"
	"net/http"
)

// Tag represents the structure of a tag.
type Tag struct {
	Name string `json:"tag_name"`
}

// Create adds a new tag.
func (t *Client) CreateTag(request Tag) ([]byte, error) {
	endpoint := "/tags/create"
	return t.request(http.MethodPost, endpoint, request)
}

// Update modifies an existing tag.
func (t *Client) UpdateTag(tagID string, request Tag) ([]byte, error) {
	endpoint := fmt.Sprintf("/tags/%s/update", tagID)
	return t.request(http.MethodPost, endpoint, request)
}

// Get retrieves details of a specific tag.
func (t *Client) RetrieveTag(tagID string) ([]byte, error) {
	endpoint := fmt.Sprintf("/tags/%s", tagID)
	return t.request(http.MethodGet, endpoint, nil)
}

// List retrieves all tags.
func (t *Client) ListTags() ([]byte, error) {
	endpoint := "/tags/list"
	return t.request(http.MethodGet, endpoint, nil)
}

// Delete removes a tag by ID.
func (t *Client) DeleteTag(tagID string) ([]byte, error) {
	endpoint := fmt.Sprintf("/tags/%s/delete", tagID)
	return t.request(http.MethodDelete, endpoint, nil)
}
