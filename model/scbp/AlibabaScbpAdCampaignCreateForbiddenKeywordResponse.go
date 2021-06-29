package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽词 API返回值 
alibaba.scbp.ad.campaign.create.forbidden.keyword

创建屏蔽词
*/
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdCampaignCreateForbiddenKeywordResponse
}

// 创建屏蔽词 成功返回结果
type AlibabaScbpAdCampaignCreateForbiddenKeywordResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_campaign_create_forbidden_keyword_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
