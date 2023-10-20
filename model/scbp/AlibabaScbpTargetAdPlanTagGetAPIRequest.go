package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplantaggetAPIRequest 定向推广-获取计划的定向溢价数据 API请求
// alibaba.scbp.target.ad.plan.tag.get
//
// 定向推广-获取计划的定向溢价数据
type AlibabascbptargetadplantaggetAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
}

// NewAlibabascbptargetadplantaggetRequest 初始化AlibabascbptargetadplantaggetAPIRequest对象
func NewAlibabascbptargetadplantaggetRequest() *AlibabascbptargetadplantaggetAPIRequest {
	return &AlibabascbptargetadplantaggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplantaggetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.tag.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplantaggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplantaggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *AlibabascbptargetadplantaggetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbptargetadplantaggetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
