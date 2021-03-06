package util

// AlibabaRetailDeviceTradeShipResult 结构体
type AlibabaRetailDeviceTradeShipResult struct {
	// errorInfos
	ErrorInfos []ErrorInfo `json:"error_infos,omitempty" xml:"error_infos>error_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
