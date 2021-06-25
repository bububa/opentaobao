package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除屏蔽词 APIRequest
alibaba.scbp.ad.campaign.delete.forbidden.keyword

删除屏蔽词
*/
type AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest struct {
    model.Params

    // 计划id
    campaignId   int64 

    // 请求参数
    forbiddenKeywordBatchOperation   *ForbiddenKeywordBatchOperationDto 

    // 用户信息
    topContext   *TopContextDto 

}

func NewAlibabaScbpAdCampaignDeleteForbiddenKeywordRequest() *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest{
    return &AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.delete.forbidden.keyword"
}

func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetCampaignId() int64 {
    return r.campaignId
}

func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) SetForbiddenKeywordBatchOperation(forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
    r.forbiddenKeywordBatchOperation = forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", forbiddenKeywordBatchOperation)
    return nil
}

func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
    return r.forbiddenKeywordBatchOperation
}

func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetTopContext() *TopContextDto {
    return r.topContext
}

