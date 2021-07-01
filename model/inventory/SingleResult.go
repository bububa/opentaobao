package inventory

// SingleResult 结构体
type SingleResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// data
	AdjustResults []InventoryCheckResultDto `json:"adjust_results,omitempty" xml:"adjust_results>inventory_check_result_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 如果是失败，可能是部分失败。如果是成功，则全部成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 地点关系
	LocationRelationList []LocationRelationDto `json:"location_relation_list,omitempty" xml:"location_relation_list>location_relation_dto,omitempty"`
}
