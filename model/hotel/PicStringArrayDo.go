package hotel

// PicStringArrayDo 结构体
type PicStringArrayDo struct {
	// prefix
	Prefix string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// sources
	Sources []string `json:"sources,omitempty" xml:"sources>string,omitempty"`
	// suffixs
	Suffixs []string `json:"suffixs,omitempty" xml:"suffixs>string,omitempty"`
}
