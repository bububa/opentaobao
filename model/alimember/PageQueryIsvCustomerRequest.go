package alimember

// PageQueryIsvCustomerRequest 结构体
type PageQueryIsvCustomerRequest struct {
	// 商户id
	OpenMerchantId string `json:"open_merchant_id,omitempty" xml:"open_merchant_id,omitempty"`
	// 第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小，最大不超过1000
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
