package icburfq

// RfqAnnexFileRemoteDto 结构体
type RfqAnnexFileRemoteDto struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 唯一文件名
	UniqueFileName string `json:"unique_file_name,omitempty" xml:"unique_file_name,omitempty"`
}
