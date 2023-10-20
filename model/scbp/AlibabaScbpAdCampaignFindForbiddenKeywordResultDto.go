package scbp

// AlibabaScbpAdCampaignFindForbiddenKeywordResultDto 结构体
type AlibabaScbpAdCampaignFindForbiddenKeywordResultDto struct {
	// 服务出参
	ResultList []AlibabaScbpAdCampaignFindForbiddenKeywordResult `json:"result_list,omitempty" xml:"result_list>alibaba_scbp_ad_campaign_find_forbidden_keyword_result,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
