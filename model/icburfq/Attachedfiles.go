package icburfq

// Attachedfiles 结构体
type Attachedfiles struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件地址
	FileUrl string `json:"file_url,omitempty" xml:"file_url,omitempty"`
}
