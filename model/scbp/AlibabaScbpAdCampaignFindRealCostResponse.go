package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划消耗数据 APIResponse
alibaba.scbp.ad.campaign.find.real.cost

批量查询计划消耗数据
*/
type AlibabaScbpAdCampaignFindRealCostAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdCampaignFindRealCostResponse `json:"alibaba_scbp_ad_campaign_find_real_cost_response,omitempty"` 
    AlibabaScbpAdCampaignFindRealCostResponse
}

/* model for simplify = false
type AlibabaScbpAdCampaignFindRealCostResponse struct {

    // 返回数据结果，json数据，key是campaignId,value是消耗数据信息
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdCampaignFindRealCostResponse struct {

    // 返回数据结果，json数据，key是campaignId,value是消耗数据信息
    Result   string `json:"result,omitempty"`

}
