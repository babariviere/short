// Code generated by ogen, DO NOT EDIT.

package oas

type CreatePostBadRequest struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *CreatePostBadRequest) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *CreatePostBadRequest) SetMessage(val OptString) {
	s.Message = val
}

func (*CreatePostBadRequest) createPostRes() {}

type CreatePostCreated struct {
	// Created shorten URL. Going to this URL should redirect to URL from request body.
	Shorten OptString `json:"shorten"`
}

// GetShorten returns the value of Shorten.
func (s *CreatePostCreated) GetShorten() OptString {
	return s.Shorten
}

// SetShorten sets the value of Shorten.
func (s *CreatePostCreated) SetShorten(val OptString) {
	s.Shorten = val
}

func (*CreatePostCreated) createPostRes() {}

type CreatePostReq struct {
	// URL to shorten.
	URL string `json:"url"`
}

// GetURL returns the value of URL.
func (s *CreatePostReq) GetURL() string {
	return s.URL
}

// SetURL sets the value of URL.
func (s *CreatePostReq) SetURL(val string) {
	s.URL = val
}

// HashGetNotFound is response for HashGet operation.
type HashGetNotFound struct{}

func (*HashGetNotFound) hashGetRes() {}

// HashGetTemporaryRedirect is response for HashGet operation.
type HashGetTemporaryRedirect struct {
	Location OptString
}

// GetLocation returns the value of Location.
func (s *HashGetTemporaryRedirect) GetLocation() OptString {
	return s.Location
}

// SetLocation sets the value of Location.
func (s *HashGetTemporaryRedirect) SetLocation(val OptString) {
	s.Location = val
}

func (*HashGetTemporaryRedirect) hashGetRes() {}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}