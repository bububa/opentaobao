package axindata

// FscDivisionApplyRequest 结构体
type FscDivisionApplyRequest struct {
	// 城市新增列表
	DivisionList []FscDivisionApplyApiDto `json:"division_list,omitempty" xml:"division_list>fsc_division_apply_api_dto,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
