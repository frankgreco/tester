package apis

// RequestDetails represents the data model that will in an http response
type RequestDetails struct {
	Method string
	Path   string
	Query  map[string]string
}
