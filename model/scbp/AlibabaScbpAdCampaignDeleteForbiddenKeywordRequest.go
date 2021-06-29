package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除屏蔽词 API请求
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

// 初始化AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest对象
func NewAlibabaScbpAdCampaignDeleteForbiddenKeywordRequest() *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest{
    return &AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.delete.forbidden.keyword"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetCampaignId() int64 {
    return r.campaignId
}
// ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) SetForbiddenKeywordBatchOperation(forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
    r.forbiddenKeywordBatchOperation = forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", forbiddenKeywordBatchOperation)
    return nil
}

// ForbiddenKeywordBatchOperation Getter
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
    return r.forbiddenKeywordBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) SetTopContext(topContext *TopContextDto) error {
    r.topContext = topContext
    r.Set("top_context", topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest) GetTopContext() *TopContextDto {
    return r.topContext
}
