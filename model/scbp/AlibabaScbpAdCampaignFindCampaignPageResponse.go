package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询计划 APIResponse
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划
*/
type AlibabaScbpAdCampaignFindCampaignPageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_campaign_find_campaign_page_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总数量
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"