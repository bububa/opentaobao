package legalcase

// RefStandpointModel 结构体
type RefStandpointModel struct {
	// 口径描述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 答辩口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 案件id
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 口径id
	StandpointId int64 `json:"standpoint_id,omitempty" xml:"standpoint_id,omitempty"`
	// 口径版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
