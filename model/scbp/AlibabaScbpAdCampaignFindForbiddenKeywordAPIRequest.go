package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaignfindforbiddenkeywordAPIRequest 查询屏蔽词 API请求
// alibaba.scbp.ad.campaign.find.forbidden.keyword
//
// 查询屏蔽词
type AlibabascbpadcampaignfindforbiddenkeywordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabascbpadcampaignfindforbiddenkeywordRequest 初始化AlibabascbpadcampaignfindforbiddenkeywordAPIRequest对象
func NewAlibabascbpadcampaignfindforbiddenkeywordRequest() *AlibabascbpadcampaignfindforbiddenkeywordAPIRequest {
	return &AlibabascbpadcampaignfindforbiddenkeywordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.forbidden.keyword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadcampaignfindforbiddenkeywordAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
