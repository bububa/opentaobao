package nrt

// NrtStoreQueryDto 结构体
type NrtStoreQueryDto struct {
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 摊位编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
