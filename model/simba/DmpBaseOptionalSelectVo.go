package simba

// DmpBaseOptionalSelectVo 结构体
type DmpBaseOptionalSelectVo struct {
	// 数据来源app
	AppIdList []int64 `json:"app_id_list,omitempty" xml:"app_id_list>int64,omitempty"`
	// 来源展示名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
}
