package alitripmerchant

// CardSystemVo 结构体
type CardSystemVo struct {
	// 会员卡详细信息
	MemberCardDetailList []MemberCardDetailVo `json:"member_card_detail_list,omitempty" xml:"member_card_detail_list>member_card_detail_vo,omitempty"`
}
