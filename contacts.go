package yonoma

import (
	"fmt"
	"net/http"
)

// Contact represents a user's contact details.
type Contact struct {
	Email  string      `json:"email"`
	Status string      `json:"status"`
	Data   ContactData `json:"data"`
}

// ContactData stores additional information about a contact.
type ContactData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"dateOfBirth"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Zipcode     string `json:"zipcode"`
}

// Status represents a contact's subscription status.
type Status struct {
	Status string `json:"status"`
}

// Client handles API operations related to Client.
// type Client struct {
// 	Client *Client
// }

// Create adds a new contact to a specified list.
func (c *Client) CreateContact(listID string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("/contacts/%s/create", listID)
	return c.request(http.MethodPost, endpoint, contact)
}

// Unsubscribe removes a contact from a list.
func (c *Client) UnsubscribeContact(listID string, contactID string, status Status) ([]byte, error) {
	endpoint := fmt.Sprintf("/contacs/%s/status/%s", listID, contactID)
	return c.request(http.MethodPost, endpoint, status)
}

// AddTag assigns a tag to a specific contact.
func (c *Client) AddTag(contactID string, tagID string) ([]byte, error) {
	endpoint := fmt.Sprintf("/contacts/tags/%s/add", contactID)
	return c.request(http.MethodPost, endpoint, tagID)
}

// RemoveTag deletes a tag from a specific contact.
func (c *Client) RemoveTag(contactID string, tagID string) ([]byte, error) {
	endpoint := fmt.Sprintf("/contacts/tags/%s/remove", contactID)
	return c.request(http.MethodDelete, endpoint, tagID)
}
