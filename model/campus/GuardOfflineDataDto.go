package campus

// GuardOfflineDataDto 结构体
type GuardOfflineDataDto struct {
	// 文件url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// guard
	Guard *GuardDto `json:"guard,omitempty" xml:"guard,omitempty"`
}
