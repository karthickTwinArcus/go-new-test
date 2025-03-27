package yonoma

import (
	"encoding/json"
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
func (c *Client) CreateContact(listID string, contact Contact) (*Contact, error) {
	endpoint := fmt.Sprintf("/contacts/%s/create", listID)

	// Send request
	response, err := c.request(http.MethodPost, endpoint, contact)
	if err != nil {
		return nil, err
	}

	// Parse response
	var createdContact Contact
	if err := json.Unmarshal(response, &createdContact); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &createdContact, nil
}

// Unsubscribe removes a contact from a list.
func (c *Client) UnsubscribeContact(listID, contactID string, status Status) error {
	endpoint := fmt.Sprintf("/contacs/%s/status/%s", listID, contactID)

	// Send request
	_, err := c.request(http.MethodPost, endpoint, status)
	return err
}

// AddTag assigns a tag to a specific contact.
func (c *Client) AddTag(contactID, tagID string) error {
	endpoint := fmt.Sprintf("/contacts/tags/%s/add", contactID)

	// Prepare request payload
	payload := map[string]string{"tag_id": tagID}

	// Send request
	_, err := c.request(http.MethodPost, endpoint, payload)
	return err
}

// RemoveTag deletes a tag from a specific contact.
func (c *Client) RemoveTag(contactID, tagID string) error {
	endpoint := fmt.Sprintf("/contacts/tags/%s/remove", contactID)

	// Send request
	_, err := c.request(http.MethodDelete, endpoint, tagID)
	return err
}
