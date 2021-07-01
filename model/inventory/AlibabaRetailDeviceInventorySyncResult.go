package inventory

// AlibabaRetailDeviceInventorySyncResult 结构体
type AlibabaRetailDeviceInventorySyncResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
