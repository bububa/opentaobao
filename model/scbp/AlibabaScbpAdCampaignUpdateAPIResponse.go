package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改计划 API返回值 
alibaba.scbp.ad.campaign.update

修改计划
*/
type AlibabaScbpAdCampaignUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignUpdateAPIResponseModel
}

// 修改计划 成功返回结果
type AlibabaScbpAdCampaignUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 修改成功数
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
