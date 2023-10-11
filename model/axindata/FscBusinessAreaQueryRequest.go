package axindata

// FscBusinessAreaQueryRequest 结构体
type FscBusinessAreaQueryRequest struct {
	// 区域类型 1:出境 2:国内
	BusinessAreaRange string `json:"business_area_range,omitempty" xml:"business_area_range,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}
