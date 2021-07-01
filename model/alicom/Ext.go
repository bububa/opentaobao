package alicom

// Ext 结构体
type Ext struct {
	// 奖品ID，传入即按指定ID兑换对应的权益
	AwardId string `json:"award_id,omitempty" xml:"award_id,omitempty"`
}
