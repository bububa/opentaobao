package tmallservice

// ServiceStoreCreateDto 结构体
type ServiceStoreCreateDto struct {
	// 秘钥--内嵌核销页面使用
	Secret string `json:"secret,omitempty" xml:"secret,omitempty"`
	// 网点id
	ServiceStoreId int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
}
