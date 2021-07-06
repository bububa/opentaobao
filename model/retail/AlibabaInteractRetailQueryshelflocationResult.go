package retail

// AlibabaInteractRetailQueryshelflocationResult 结构体
type AlibabaInteractRetailQueryshelflocationResult struct {
	// warningInfos
	WarningInfos []string `json:"warning_infos,omitempty" xml:"warning_infos>string,omitempty"`
	// errorInfos
	ErrorInfos []string `json:"error_infos,omitempty" xml:"error_infos>string,omitempty"`
	// module
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
