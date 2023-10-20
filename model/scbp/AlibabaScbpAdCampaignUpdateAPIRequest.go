package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdCampaignUpdateAPIRequest 修改计划 API请求
// alibaba.scbp.ad.campaign.update
//
// 修改计划
type AlibabaScbpAdCampaignUpdateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 修改数据
	_campaignOperation *CampaignOperationDto
}

// NewAlibabaScbpAdCampaignUpdateRequest 初始化AlibabaScbpAdCampaignUpdateAPIRequest对象
func NewAlibabaScbpAdCampaignUpdateRequest() *AlibabaScbpAdCampaignUpdateAPIRequest {
	return &AlibabaScbpAdCampaignUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdCampaignUpdateAPIRequest) Reset() {
	r._topContext = nil
	r._campaignOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdCampaignUpdateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignOperation is CampaignOperation Setter
// 修改数据
func (r *AlibabaScbpAdCampaignUpdateAPIRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDto) error {
	r._campaignOperation = _campaignOperation
	r.Set("campaign_operation", _campaignOperation)
	return nil
}

// GetCampaignOperation CampaignOperation Getter
func (r AlibabaScbpAdCampaignUpdateAPIRequest) GetCampaignOperation() *CampaignOperationDto {
	return r._campaignOperation
}

var poolAlibabaScbpAdCampaignUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdCampaignUpdateRequest()
	},
}

// GetAlibabaScbpAdCampaignUpdateRequest 从 sync.Pool 获取 AlibabaScbpAdCampaignUpdateAPIRequest
func GetAlibabaScbpAdCampaignUpdateAPIRequest() *AlibabaScbpAdCampaignUpdateAPIRequest {
	return poolAlibabaScbpAdCampaignUpdateAPIRequest.Get().(*AlibabaScbpAdCampaignUpdateAPIRequest)
}

// ReleaseAlibabaScbpAdCampaignUpdateAPIRequest 将 AlibabaScbpAdCampaignUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdCampaignUpdateAPIRequest(v *AlibabaScbpAdCampaignUpdateAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdCampaignUpdateAPIRequest.Put(v)
}
