// Code generated by ogen, DO NOT EDIT.

package oas

type CreateShortURLBadRequest struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *CreateShortURLBadRequest) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *CreateShortURLBadRequest) SetMessage(val OptString) {
	s.Message = val
}

func (*CreateShortURLBadRequest) createShortURLRes() {}

type CreateShortURLCreated struct {
	// Created shorten URL. Going to this URL should redirect to URL from request body.
	Shorten OptString `json:"shorten"`
}

// GetShorten returns the value of Shorten.
func (s *CreateShortURLCreated) GetShorten() OptString {
	return s.Shorten
}

// SetShorten sets the value of Shorten.
func (s *CreateShortURLCreated) SetShorten(val OptString) {
	s.Shorten = val
}

func (*CreateShortURLCreated) createShortURLRes() {}

type CreateShortURLReq struct {
	// URL to shorten.
	URL string `json:"url"`
}

// GetURL returns the value of URL.
func (s *CreateShortURLReq) GetURL() string {
	return s.URL
}

// SetURL sets the value of URL.
func (s *CreateShortURLReq) SetURL(val string) {
	s.URL = val
}

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

// RedirectLongURLNotFound is response for RedirectLongURL operation.
type RedirectLongURLNotFound struct{}

func (*RedirectLongURLNotFound) redirectLongURLRes() {}

// RedirectLongURLTemporaryRedirect is response for RedirectLongURL operation.
type RedirectLongURLTemporaryRedirect struct {
	Location OptString
}

// GetLocation returns the value of Location.
func (s *RedirectLongURLTemporaryRedirect) GetLocation() OptString {
	return s.Location
}

// SetLocation sets the value of Location.
func (s *RedirectLongURLTemporaryRedirect) SetLocation(val OptString) {
	s.Location = val
}

func (*RedirectLongURLTemporaryRedirect) redirectLongURLRes() {}
