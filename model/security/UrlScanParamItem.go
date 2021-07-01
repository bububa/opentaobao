package security

// UrlScanParamItem 结构体
type UrlScanParamItem struct {
	// 需要扫描的url
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 参数标记，用于识别返回结果对应的参数
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
}
