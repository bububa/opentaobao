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
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest struct {
    model.Params
    // 请求参数
    _forbiddenKeywordBatchOperation   *ForbiddenKeywordBatchOperationDto
    // 计划id
    _campaignId   int64
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest对象
func NewAlibabaScbpAdCampaignCreateForbiddenKeywordRequest() *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest{
    return &AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.create.forbidden.keyword"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) SetForbiddenKeywordBatchOperation(_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
    r._forbiddenKeywordBatchOperation = _forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", _forbiddenKeywordBatchOperation)
    return nil
}

// ForbiddenKeywordBatchOperation Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
    return r._forbiddenKeywordBatchOperation
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
