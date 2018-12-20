package types

// Book data structure for ginkgo testing
type Book struct {
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	PageCount int    `json:"page_count"`
}

// Books data structure for ginkgo testing
type Books struct {
	Items []Book `json:"items"`
}
