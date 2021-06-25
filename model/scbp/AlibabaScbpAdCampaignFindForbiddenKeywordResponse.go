package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询屏蔽词 APIResponse
alibaba.scbp.ad.campaign.find.forbidden.keyword

查询屏蔽词
*/
type AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdCampaignFindForbiddenKeywordResponse `json:"alibaba_scbp_ad_campaign_find_forbidden_keyword_response,omitempty"`
}

type AlibabaScbpAdCampaignFindForbiddenKeywordResponse struct {

    // 返回数据
    Result   *AlibabaScbpAdCampaignFindForbiddenKeywordResultDto `json:"result,omitempty"`

}