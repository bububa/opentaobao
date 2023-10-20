package scbp

// AlibabascbpadcampaignfindforbiddenkeywordResultDto 结构体
type AlibabascbpadcampaignfindforbiddenkeywordResultDto struct {
	// 服务出参
	ResultList []AlibabascbpadcampaignfindforbiddenkeywordResult `json:"result_list,omitempty" xml:"result_list>alibabascbpadcampaignfindforbiddenkeyword_result,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
