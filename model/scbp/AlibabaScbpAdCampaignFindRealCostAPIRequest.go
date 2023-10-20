package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaignfindrealcostAPIRequest 批量查询计划消耗数据 API请求
// alibaba.scbp.ad.campaign.find.real.cost
//
// 批量查询计划消耗数据
type AlibabascbpadcampaignfindrealcostAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 系统自动生成
	_campaignQuery *CampaignQueryDto
}

// NewAlibabascbpadcampaignfindrealcostRequest 初始化AlibabascbpadcampaignfindrealcostAPIRequest对象
func NewAlibabascbpadcampaignfindrealcostRequest() *AlibabascbpadcampaignfindrealcostAPIRequest {
	return &AlibabascbpadcampaignfindrealcostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaignfindrealcostAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.real.cost"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaignfindrealcostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaignfindrealcostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaignfindrealcostAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaignfindrealcostAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignQuery is CampaignQuery Setter
// 系统自动生成
func (r *AlibabascbpadcampaignfindrealcostAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r AlibabascbpadcampaignfindrealcostAPIRequest) GetCampaignQuery() *CampaignQueryDto {
	return r._campaignQuery
}
