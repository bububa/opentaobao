package flight

// PenaltyItemDto 结构体
type PenaltyItemDto struct {
	// 退改签规则集合
	PenaltyBOList []PenaltyDto `json:"penalty_b_o_list,omitempty" xml:"penalty_b_o_list>penalty_dto,omitempty"`
}
