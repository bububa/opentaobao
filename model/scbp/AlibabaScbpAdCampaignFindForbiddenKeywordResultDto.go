package scbp

// AlibabaScbpAdCampaignFindForbiddenKeywordResultDto 
/* model for simplify = false
type AlibabaScbpAdCampaignFindForbiddenKeywordResultDto struct {

    // 信息
    
    Msg   string `json:"msg,omitempty"`
    

    // 返回code
    
    Code   string `json:"code,omitempty"`
    

    // 执行结果
    
    Success   bool `json:"success,omitempty"`
    

    // 服务出参
    
    ResultList  struct {
        AlibabaScbpAdCampaignFindForbiddenKeywordResult  []AlibabaScbpAdCampaignFindForbiddenKeywordResult `json:"alibaba_scbp_ad_campaign_find_forbidden_keyword_result,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

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
