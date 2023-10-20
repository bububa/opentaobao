package ascpchannel

// AlibabaascpchanneldistributorproductlistData 结构体
type AlibabaascpchanneldistributorproductlistData struct {
	// 产品列表
	Products []Products `json:"products,omitempty" xml:"products>products,omitempty"`
	// 返回值总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
