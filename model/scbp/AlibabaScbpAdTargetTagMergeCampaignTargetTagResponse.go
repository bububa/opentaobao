package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标签增删改 APIResponse
alibaba.scbp.ad.target.tag.merge.campaign.target.tag

标签增删改
*/
type AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse
}

type AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_merge_campaign_target_tag_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
