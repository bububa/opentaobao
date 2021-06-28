package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除计划 APIResponse
alibaba.scbp.ad.campaign.delete

删除计划
*/
type AlibabaScbpAdCampaignDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignDeleteResponse
}

type AlibabaScbpAdCampaignDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除成功条数
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
