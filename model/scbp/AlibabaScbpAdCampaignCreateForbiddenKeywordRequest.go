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
    _forbiddenKeywordBatchOperation   *ForbiddenKeywordBatchOperationDTO
    // 计划id
    _campaignId   int64
    // 用户信息
    _topContext   *TopContextDTO
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
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetForbiddenKeywordBatchOperation(_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDTO) error {
    r._forbiddenKeywordBatchOperation = _forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", _forbiddenKeywordBatchOperation)
    return nil
}

// ForbiddenKeywordBatchOperation Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDTO {
    return r._forbiddenKeywordBatchOperation
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetCampaignId() int64 {
    return r._campaignId
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) SetTopContext(_topContext *TopContextDTO) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordRequest) GetTopContext() *TopContextDTO {
    return r._topContext
}
