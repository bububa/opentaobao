package btrip

// OpenUserAffiliateDo 结构体
type OpenUserAffiliateDo struct {
	// 出行人ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 出行人名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}
