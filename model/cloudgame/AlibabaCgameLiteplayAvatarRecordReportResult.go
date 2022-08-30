package cloudgame

// AlibabaCgameLiteplayAvatarRecordReportResult 结构体
type AlibabaCgameLiteplayAvatarRecordReportResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// str消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回消息体
	Data *TopRecordCallbackResp `json:"data,omitempty" xml:"data,omitempty"`
}
