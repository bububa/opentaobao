package ascpchannel

// AlibabaascpglobalsupplieritemlistinfoqueryData 结构体
type AlibabaascpglobalsupplieritemlistinfoqueryData struct {
	// 返回货品list
	PageData []PageData `json:"page_data,omitempty" xml:"page_data>page_data,omitempty"`
	// 每页条数，最大不超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 货品总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
