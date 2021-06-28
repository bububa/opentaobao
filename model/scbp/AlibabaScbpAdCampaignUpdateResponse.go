package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改计划 APIResponse
alibaba.scbp.ad.campaign.update

修改计划
*/
type AlibabaScbpAdCampaignUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_campaign_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改成功数
    
    Result   int64 `json:"result,omitempty" xml:"