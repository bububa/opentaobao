package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改计划 APIResponse
alibaba.scbp.ad.campaign.update

修改计划
*/
type AlibabaScbpAdCampaignUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdCampaignUpdateResponse `json:"alibaba_scbp_ad_campaign_update_response,omitempty"`
}

type AlibabaScbpAdCampaignUpdateResponse struct {

    // 修改成功数
    Result   int64 `json:"result,omitempty"`

}
