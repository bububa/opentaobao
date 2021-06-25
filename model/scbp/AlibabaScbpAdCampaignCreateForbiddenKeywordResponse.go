package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽词 APIResponse
alibaba.scbp.ad.campaign.create.forbidden.keyword

创建屏蔽词
*/
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdCampaignCreateForbiddenKeywordResponse `json:"alibaba_scbp_ad_campaign_create_forbidden_keyword_response,omitempty"`
}

type AlibabaScbpAdCampaignCreateForbiddenKeywordResponse struct {

    // 返回结果
    Result   string `json:"result,omitempty"`

}
