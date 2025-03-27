## [Yonoma](https://yonoma.io/) Go Email Marketing SDK
### Install
```bash
go get github.com/YonomaHQ/yonoma
```
### Setup
First, you need to get an API key:
```go
client := yonoma.NewClient("api-key") 
```
### Usage
## Send your email:
```go
email := yonoma.Email{
    FromEmail:    "updates@yonoma.io",
    ToEmail:      "email@yourdomain.com",
    Subject:      "Welcome to Yonoma - You're In!",
    MailTemplate: "We're excited to welcome you to Yonoma! Your successful signup marks the beginning of what we hope will be an exceptional journey."
}
response, err := client.Send(email)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
## Contacts
#### Create new contact
```go
contactData := yonoma.Contact{
    Email:  "glennjacob@example.com",
    Status: "Subscribed" | "Unsubscribed",
    FirstName: string,
    LastName:  string,
    Phone:     string,
    Gender:    string,
    Address:   string,
    City:      string,
    State:     string,
    Country:   string,
    Zipcode:   string,
}
responseC, err := client.CreateContact("List Id", contactData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(responseC))
}
```
#### Update contact
```go
contactData := yonoma.Status{
	Status: "Subscribed" | "Unsubscribed",
}
response, err := client.UnsubscribeContact("List Id", "Contact Id", contactData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### Add tag to contact
```go
contactData := yonoma.TagId{
    TagId: "Tag Id",
}
response, err := client.AddContactTag("Contact Id", contactData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### Remove tag from contact
```go
contactData := yonoma.TagId{
    TagId: "Tag Id",
}
response, err := client.RemoveContactTag("Contact Id", contactData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
## Managing Tags
#### Create a Tag
```go
tagData := yonoma.Tag{
	Name: "Tag Name",
}
response, err := client.CreateTag(tagData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### Update a Tag
```go
tagData := yonoma.Tag{
	Name: "Tag Name",
}
response, err := client.UpdateTag(tagData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### Delete a Tag
```go
response, err := client.DeleteTag("Tag Id")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### Retrieve a Specific Tag
```go
response, err := client.RetrieveTag("Tag Id")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### List All Tags
```go
response, err := client.ListTags()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
## Managing Lists
#### Create a List
```go
listData := yonoma.List{
	Name: "List Name",
}
response, err := client.CreateList(listData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}

```
#### Update a List
```go
listData := yonoma.List{
	Name: "List Name",
}
response, err := client.UpdateList("List Id", listData)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}

```
#### Delete a List
```go
response, err := client.DeleteList("List Id")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### Retrieve a Specific List
```go
response, err := client.RetrieveList("List Id")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```
#### List All Lists
```go
response, err := client.ListLists()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Response:", string(response))
}
```

