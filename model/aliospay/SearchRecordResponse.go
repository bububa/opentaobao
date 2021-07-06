package aliospay

// SearchRecordResponse 结构体
type SearchRecordResponse struct {
	// 支付记录列表
	Datas []PayRecordData `json:"datas,omitempty" xml:"datas>pay_record_data,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
