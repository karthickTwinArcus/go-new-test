package yonoma

import (
	"fmt"
	"net/http"
)

// ListRequest represents the structure for creating or updating a list.
type List struct {
	Name string `json:"list_name"`
}

// Create adds a new list.
func (l *Client) CreateList(request List) ([]byte, error) {
	endpoint := "/lists/create"
	return l.request(http.MethodPost, endpoint, request)
}

// Update modifies an existing mailing list.
func (l *Client) UpdateList(listID string, request List) ([]byte, error) {
	endpoint := fmt.Sprintf("/lists/%s/update", listID)
	return l.request(http.MethodPost, endpoint, request)
}

// Get retrieves details of a specific list.
func (l *Client) RetrieveList(listID string) ([]byte, error) {
	endpoint := fmt.Sprintf("/lists/%s", listID)
	return l.request(http.MethodGet, endpoint, nil)
}

// List retrieves all mailing lists.
func (l *Client) ListLists() ([]byte, error) {
	endpoint := "/lists/list"
	return l.request(http.MethodGet, endpoint, nil)
}

// Delete removes a mailing list by ID.
func (l *Client) DeleteList(listID string) ([]byte, error) {
	endpoint := fmt.Sprintf("/lists/%s/delete", listID)
	return l.request(http.MethodDelete, endpoint, nil)
}
