package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除屏蔽词 APIResponse
alibaba.scbp.ad.campaign.delete.forbidden.keyword

删除屏蔽词
*/
type AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdCampaignDeleteForbiddenKeywordResponse `json:"alibaba_scbp_ad_campaign_delete_forbidden_keyword_response,omitempty"` 
    AlibabaScbpAdCampaignDeleteForbiddenKeywordResponse
}

/* model for simplify = false
type AlibabaScbpAdCampaignDeleteForbiddenKeywordResponse struct {

    // 返回结果
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdCampaignDeleteForbiddenKeywordResponse struct {

    // 返回结果
    Result   int64 `json:"result,omitempty"`

}
