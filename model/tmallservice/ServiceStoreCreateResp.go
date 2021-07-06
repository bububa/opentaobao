package tmallservice

// ServiceStoreCreateResp 结构体
type ServiceStoreCreateResp struct {
	// 门店密钥
	Secret string `json:"secret,omitempty" xml:"secret,omitempty"`
	// 门店id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
