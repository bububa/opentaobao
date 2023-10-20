package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupfindforbiddenproductAPIRequest 查询屏蔽品 API请求
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
type AlibabascbpadgroupfindforbiddenproductAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabascbpadgroupfindforbiddenproductRequest 初始化AlibabascbpadgroupfindforbiddenproductAPIRequest对象
func NewAlibabascbpadgroupfindforbiddenproductRequest() *AlibabascbpadgroupfindforbiddenproductAPIRequest {
	return &AlibabascbpadgroupfindforbiddenproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadgroupfindforbiddenproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.group.find.forbidden.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadgroupfindforbiddenproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadgroupfindforbiddenproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadgroupfindforbiddenproductAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadgroupfindforbiddenproductAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadgroupfindforbiddenproductAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadgroupfindforbiddenproductAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
