package security

// AppInfoBatchItem 结构体
type AppInfoBatchItem struct {
	// 待扫描的应用信息dataType=3时填应用的md5 dataType=4时填应用的url
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 应用的md5，dataType＝4时必填
	Md5 string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 应用的大小(单位byte)，dataType＝4时必填
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}
