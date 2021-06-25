package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 APIResponse
alibaba.scbp.ad.campaign.delete

删除计划
*/
type AlibabaScbpAdCampaignDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdCampaignDeleteResponse `json:"alibaba_scbp_ad_campaign_delete_response,omitempty"`
}

type AlibabaScbpAdCampaignDeleteResponse struct {

    // 删除成功条数
    Result   int64 `json:"result,omitempty"`

}
