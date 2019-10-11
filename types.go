package main

// Label contains information of a GitHub label.
type Label struct {
	Name        string
	Color       string
	Description string
}

// PageInfo contains pagination metadata of a GraphQL query.
type PageInfo struct {
	EndCursor   string
	HasNextPage bool
}
