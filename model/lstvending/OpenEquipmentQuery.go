package lstvending

// OpenEquipmentQuery 结构体
type OpenEquipmentQuery struct {
	// 排序条件
	SortParamList []SortParam `json:"sort_param_list,omitempty" xml:"sort_param_list>sort_param,omitempty"`
	// 修改时间
	GmtModifiedRange *Range `json:"gmt_modified_range,omitempty" xml:"gmt_modified_range,omitempty"`
	// 每页记录数
	PageRows int64 `json:"page_rows,omitempty" xml:"page_rows,omitempty"`
	// 页数
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 创建时间
	GmtCreateRange *Range `json:"gmt_create_range,omitempty" xml:"gmt_create_range,omitempty"`
}
