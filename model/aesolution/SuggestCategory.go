package aesolution

// SuggestCategory 结构体
type SuggestCategory struct {
	// category id path
	IdPath []int64 `json:"id_path,omitempty" xml:"id_path>int64,omitempty"`
	// category name path
	NamePath []string `json:"name_path,omitempty" xml:"name_path>string,omitempty"`
	// category name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// category id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
