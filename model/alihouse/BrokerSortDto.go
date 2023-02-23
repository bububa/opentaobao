package alihouse

// BrokerSortDto 结构体
type BrokerSortDto struct {
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 排序号
	SortNo int64 `json:"sort_no,omitempty" xml:"sort_no,omitempty"`
}
