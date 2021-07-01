package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 API返回值 
alibaba.scbp.ad.campaign.delete

删除计划
*/
type AlibabaScbpAdCampaignDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignDeleteAPIResponseModel
}

// 删除计划 成功返回结果
type AlibabaScbpAdCampaignDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除成功条数
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
