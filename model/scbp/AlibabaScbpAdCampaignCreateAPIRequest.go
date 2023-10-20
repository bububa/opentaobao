package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignCreateAPIRequest 创建计划 API请求
// alibaba.scbp.ad.campaign.create
//
// 创建计划
type AlibabaScbpAdCampaignCreateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 返回数据
	_campaignOperation *CampaignOperationDto
}

// NewAlibabaScbpAdCampaignCreateRequest 初始化AlibabaScbpAdCampaignCreateAPIRequest对象
func NewAlibabaScbpAdCampaignCreateRequest() *AlibabaScbpAdCampaignCreateAPIRequest {
	return &AlibabaScbpAdCampaignCreateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignCreateAPIRequest) Reset() {
	r._topContext = nil
	r._campaignOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignCreateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignOperation is CampaignOperation Setter
// 返回数据
func (r *AlibabaScbpAdCampaignCreateAPIRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDto) error {
	r._campaignOperation = _campaignOperation
	r.Set("campaign_operation", _campaignOperation)
	return nil
}

// GetCampaignOperation CampaignOperation Getter
func (r AlibabaScbpAdCampaignCreateAPIRequest) GetCampaignOperation() *CampaignOperationDto {
	return r._campaignOperation
}

var poolAlibabaScbpAdCampaignCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignCreateRequest()
	},
}

// GetAlibabaScbpAdCampaignCreateRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignCreateAPIRequest
func GetAlibabaScbpAdCampaignCreateAPIRequest() *AlibabaScbpAdCampaignCreateAPIRequest {
	return poolAlibabaScbpAdCampaignCreateAPIRequest.Get().(*AlibabaScbpAdCampaignCreateAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignCreateAPIRequest 将 AlibabaScbpAdCampaignCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignCreateAPIRequest(v *AlibabaScbpAdCampaignCreateAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignCreateAPIRequest.Put(v)
}
