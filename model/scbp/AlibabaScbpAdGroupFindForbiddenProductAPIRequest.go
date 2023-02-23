package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupFindForbiddenProductAPIRequest 查询屏蔽品 API请求
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
type AlibabaScbpAdGroupFindForbiddenProductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabaScbpAdGroupFindForbiddenProductRequest 初始化AlibabaScbpAdGroupFindForbiddenProductAPIRequest对象
func NewAlibabaScbpAdGroupFindForbiddenProductRequest() *AlibabaScbpAdGroupFindForbiddenProductAPIRequest {
	return &AlibabaScbpAdGroupFindForbiddenProductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdGroupFindForbiddenProductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.find.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdGroupFindForbiddenProductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdGroupFindForbiddenProductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdGroupFindForbiddenProductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdGroupFindForbiddenProductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabaScbpAdGroupFindForbiddenProductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabaScbpAdGroupFindForbiddenProductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
