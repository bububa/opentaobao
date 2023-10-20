package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaignfindcampaignpageAPIRequest 分页查询计划 API请求
// alibaba.scbp.ad.campaign.find.campaign.page
//
// 分页查询计划
type AlibabascbpadcampaignfindcampaignpageAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求实体类
	_campaignQuery *CampaignQueryDto
}

// NewAlibabascbpadcampaignfindcampaignpageRequest 初始化AlibabascbpadcampaignfindcampaignpageAPIRequest对象
func NewAlibabascbpadcampaignfindcampaignpageRequest() *AlibabascbpadcampaignfindcampaignpageAPIRequest {
	return &AlibabascbpadcampaignfindcampaignpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaignfindcampaignpageAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.find.campaign.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaignfindcampaignpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaignfindcampaignpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaignfindcampaignpageAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaignfindcampaignpageAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignQuery is CampaignQuery Setter
// 请求实体类
func (r *AlibabascbpadcampaignfindcampaignpageAPIRequest) SetCampaignQuery(_campaignQuery *CampaignQueryDto) error {
	r._campaignQuery = _campaignQuery
	r.Set("campaign_query", _campaignQuery)
	return nil
}

// GetCampaignQuery CampaignQuery Getter
func (r AlibabascbpadcampaignfindcampaignpageAPIRequest) GetCampaignQuery() *CampaignQueryDto {
	return r._campaignQuery
}
