package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建计划 API返回值 
alibaba.scbp.ad.campaign.create

创建计划
*/
type AlibabaScbpAdCampaignCreateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignCreateResponse
}

// 创建计划 成功返回结果
type AlibabaScbpAdCampaignCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 计划id
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
