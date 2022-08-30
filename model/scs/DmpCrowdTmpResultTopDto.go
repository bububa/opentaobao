package scs

// DmpCrowdTmpResultTopDto 结构体
type DmpCrowdTmpResultTopDto struct {
	// group_ids
	GroupIdList []int64 `json:"group_id_list,omitempty" xml:"group_id_list>int64,omitempty"`
	// highlight
	Highlight string `json:"highlight,omitempty" xml:"highlight,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 过期时间
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
