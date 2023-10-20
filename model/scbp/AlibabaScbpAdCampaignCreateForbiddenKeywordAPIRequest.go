package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest 创建屏蔽词 API请求
// alibaba.scbp.ad.campaign.create.forbidden.keyword
//
// 创建屏蔽词
type AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto
}

// NewAlibabaScbpAdCampaignCreateForbiddenKeywordRequest 初始化AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest对象
func NewAlibabaScbpAdCampaignCreateForbiddenKeywordRequest() *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest {
	return &AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) Reset() {
	r._topContext = nil
	r._campaignId = 0
	r._forbiddenKeywordBatchOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.create.forbidden.keyword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenKeywordBatchOperation is ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) SetForbiddenKeywordBatchOperation(_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
	r._forbiddenKeywordBatchOperation = _forbiddenKeywordBatchOperation
	r.Set("forbidden_keyword_batch_operation", _forbiddenKeywordBatchOperation)
	return nil
}

// GetForbiddenKeywordBatchOperation ForbiddenKeywordBatchOperation Getter
func (r AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
	return r._forbiddenKeywordBatchOperation
}

var poolAlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignCreateForbiddenKeywordRequest()
	},
}

// GetAlibabaScbpAdCampaignCreateForbiddenKeywordRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest
func GetAlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest() *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest {
	return poolAlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest.Get().(*AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest 将 AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest(v *AlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignCreateForbiddenKeywordAPIRequest.Put(v)
}
