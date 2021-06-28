package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询计划消耗数据 APIResponse
alibaba.scbp.ad.campaign.find.real.cost

批量查询计划消耗数据
*/
type AlibabaScbpAdCampaignFindRealCostAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignFindRealCostResponse
}

type AlibabaScbpAdCampaignFindRealCostResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_real_cost_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回数据结果，json数据，key是campaignId,value是消耗数据信息
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
