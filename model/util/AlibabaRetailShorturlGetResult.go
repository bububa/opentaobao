package util

// AlibabaretailshorturlgetResult 结构体
type AlibabaretailshorturlgetResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// module
	Module *ShortUrlDto `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
