package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建计划 APIResponse
alibaba.scbp.ad.campaign.create

创建计划
*/
type AlibabaScbpAdCampaignCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdCampaignCreateResponse `json:"alibaba_scbp_ad_campaign_create_response,omitempty"` 
    AlibabaScbpAdCampaignCreateResponse
}

/* model for simplify = false
type AlibabaScbpAdCampaignCreateResponse struct {

    // 计划id
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdCampaignCreateResponse struct {

    // 计划id
    Result   int64 `json:"result,omitempty"`

}
