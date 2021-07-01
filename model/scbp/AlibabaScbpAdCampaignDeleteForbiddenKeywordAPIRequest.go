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
type AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest struct {
    model.Params
    // 计划id
    _campaignId   int64
    // 请求参数
    _forbiddenKeywordBatchOperation   *ForbiddenKeywordBatchOperationDto
    // 用户信息
    _topContext   *TopContextDto
}

// 初始化AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest对象
func NewAlibabaScbpAdCampaignDeleteForbiddenKeywordRequest() *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest{
    return &AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.ad.campaign.delete.forbidden.keyword"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) SetForbiddenKeywordBatchOperation(_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
    r._forbiddenKeywordBatchOperation = _forbiddenKeywordBatchOperation
    r.Set("forbidden_keyword_batch_operation", _forbiddenKeywordBatchOperation)
    return nil
}

// ForbiddenKeywordBatchOperation Getter
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
    return r._forbiddenKeywordBatchOperation
}
// TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
    r._topContext = _topContext
    r.Set("top_context", _topContext)
    return nil
}

// TopContext Getter
func (r AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIRequest) GetTopContext() *TopContextDto {
    return r._topContext
}
