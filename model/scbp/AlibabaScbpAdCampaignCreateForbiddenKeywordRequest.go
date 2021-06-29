package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建屏蔽词 API请求
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

// 初始化AlibabaScbpAdCampaignCreateForbiddenKeywordRequest对象
func NewAlibabaScbpAdCampaignCreateForbiddenKeywordRequest() *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest{
    return &AlibabaScbpAdCampaignCreateForbiddenKeywordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.create.forbidden.keyword"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetForbiddenKeywordBatchOperation(forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
    r.forbiddenKeywordBatchOperation = forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", forbiddenKeywordBatchOperation)
    return nil
}

// ForbiddenKeywordBatchOperation Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
    return r.forbiddenKeywordBatchOperation
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetCampaignId() int64 {
    return r.campaignId
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
