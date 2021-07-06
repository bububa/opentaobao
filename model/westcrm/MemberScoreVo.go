package westcrm

// MemberScoreVo 结构体
type MemberScoreVo struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 积分
	Score int64 `json:"score,omitempty" xml:"score,omitempty"`
}
