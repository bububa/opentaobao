package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽词 APIRequest
alibaba.scbp.ad.campaign.create.forbidden.keyword

创建屏蔽词
*/
type AlibabaScbpAdCampaignCreateForbiddenKeywordRequest struct {
    model.Params

    // 请求参数
    forbiddenKeywordBatchOperation   *ForbiddenKeywordBatchOperationDto 

    // 计划id
    campaignId   int64 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdCampaignCreateForbiddenKeywordRequest() *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest{
    return &AlibabaScbpAdCampaignCreateForbiddenKeywordRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.create.forbidden.keyword"
}

func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetForbiddenKeywordBatchOperation(forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
    r.forbiddenKeywordBatchOperation = forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", forbiddenKeywordBatchOperation)
    return nil
}

func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
    return r.forbiddenKeywordBatchOperation
}

func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

