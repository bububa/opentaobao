package bus

// TvmInsuranceInfo 结构体
type TvmInsuranceInfo struct {
	// 保险名称
	InsureName string `json:"insure_name,omitempty" xml:"insure_name,omitempty"`
	// 保险唯一id
	ProductNo string `json:"product_no,omitempty" xml:"product_no,omitempty"`
	// 保险商品单价(单位分)
	InsurePrice int64 `json:"insure_price,omitempty" xml:"insure_price,omitempty"`
	// 保险状态: -1下单失败 0初始化 1已取消 2已关闭 3已挂起 4已挂起  5未知状态  6未生效 7保障中  8已失效  9退保中 10已退保 11未生效或保障中
	InsureStatus int64 `json:"insure_status,omitempty" xml:"insure_status,omitempty"`
}
