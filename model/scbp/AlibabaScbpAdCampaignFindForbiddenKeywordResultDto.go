package scbp

// AlibabaScbpAdCampaignFindForbiddenKeywordResultDto 
type AlibabaScbpAdCampaignFindForbiddenKeywordResultDto struct {

    // 信息
    Msg   string `json:"msg,omitempty"`

    // 返回code
    Code   string `json:"code,omitempty"`

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 服务出参
    ResultList   []AlibabaScbpAdCampaignFindForbiddenKeywordResult `json:"result_list,omitempty"`

}
