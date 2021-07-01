package opentrade

// RopGroupMemberInfo 结构体
type RopGroupMemberInfo struct {
	// 用户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 用户头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 1-未付款，2-已付款，4-已退款，6-交易成功，8-交易关闭
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}
