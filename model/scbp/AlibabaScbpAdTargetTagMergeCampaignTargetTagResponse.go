package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标签增删改 APIResponse
alibaba.scbp.ad.target.tag.merge.campaign.target.tag

标签增删改
*/
type AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse `json:"alibaba_scbp_ad_target_tag_merge_campaign_target_tag_response,omitempty"` 
    AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse
}

/* model for simplify = false
type AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse struct {

    // 返回值
    
    Result   int64 `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse struct {

    // 返回值
    Result   int64 `json:"result,omitempty"`

}
