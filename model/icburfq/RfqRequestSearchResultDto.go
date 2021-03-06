package icburfq

// RfqRequestSearchResultDto 结构体
type RfqRequestSearchResultDto struct {
	// RFQ列表
	RequestList []Requestlist `json:"request_list,omitempty" xml:"request_list>requestlist,omitempty"`
	// 类目列表
	CategoryList []Categorylist `json:"category_list,omitempty" xml:"category_list>categorylist,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
