package qimen

// TaobaoQimenExpressinfoQueryResponse 结构体
type TaobaoQimenExpressinfoQueryResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	ExpressInfos []ExpressInfo `json:"expressInfos,omitempty" xml:"expressInfos>express_info,omitempty"`
}
