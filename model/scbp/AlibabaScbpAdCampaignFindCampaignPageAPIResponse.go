package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询计划 API返回值 
alibaba.scbp.ad.campaign.find.campaign.page

分页查询计划
*/
type AlibabaScbpAdCampaignFindCampaignPageAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel
}

// 分页查询计划 成功返回结果
type AlibabaScbpAdCampaignFindCampaignPageAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_campaign_page_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 总数量
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 返回数据
    ResultList   []CampaignDto `json:"result_list,omitempty" xml:"result_list>campaign_dto,omitempty"`
}
