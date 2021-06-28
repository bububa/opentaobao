package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询计划 APIResponse
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划
*/
type AlibabaScbpAdCampaignFindCampaignPageAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdCampaignFindCampaignPageResponse `json:"alibaba_scbp_ad_campaign_find_campaign_page_response,omitempty"` 
    AlibabaScbpAdCampaignFindCampaignPageResponse
}

/* model for simplify = false
type AlibabaScbpAdCampaignFindCampaignPageResponse struct {

    // 总数量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 返回数据
    
    ResultList  struct {
        CampaignDto  []CampaignDto `json:"campaign_dto,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaScbpAdCampaignFindCampaignPageResponse struct {

    // 总数量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 返回数据
    ResultList   []CampaignDto `json:"result_list,omitempty"`

}
