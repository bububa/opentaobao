package cloudgame

// TopRecordCallbackResp 结构体
type TopRecordCallbackResp struct {
	// 返回消息result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// top请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 云游戏mixUserId
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// str消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
