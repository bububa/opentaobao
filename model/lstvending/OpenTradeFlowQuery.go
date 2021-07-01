package lstvending

// OpenTradeFlowQuery 结构体
type OpenTradeFlowQuery struct {
	// 修改时间
	GmtModifiedRange *Range `json:"gmt_modified_range,omitempty" xml:"gmt_modified_range,omitempty"`
	// 每页记录数
	PageRows int64 `json:"page_rows,omitempty" xml:"page_rows,omitempty"`
	// 排序条件
	SortParamList []SortParam `json:"sort_param_list,omitempty" xml:"sort_param_list>sort_param,omitempty"`
	// 设备ID
	EquipmentId int64 `json:"equipment_id,omitempty" xml:"equipment_id,omitempty"`
	// 创建时间
	GmtCreateRange *Range `json:"gmt_create_range,omitempty" xml:"gmt_create_range,omitempty"`
	// 页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
