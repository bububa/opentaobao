package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽词 API返回值 
alibaba.scbp.ad.campaign.find.forbidden.keyword

查询屏蔽词
*/
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel
}

// 查询屏蔽词 成功返回结果
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_find_forbidden_keyword_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Result   *AlibabaScbpAdCampaignFindForbiddenKeywordResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
