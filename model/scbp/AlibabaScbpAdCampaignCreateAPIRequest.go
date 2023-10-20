package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaigncreateAPIRequest 创建计划 API请求
// alibaba.scbp.ad.campaign.create
//
// 创建计划
type AlibabascbpadcampaigncreateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 返回数据
	_campaignOperation *CampaignOperationDto
}

// NewAlibabascbpadcampaigncreateRequest 初始化AlibabascbpadcampaigncreateAPIRequest对象
func NewAlibabascbpadcampaigncreateRequest() *AlibabascbpadcampaigncreateAPIRequest {
	return &AlibabascbpadcampaigncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaigncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaigncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaigncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaigncreateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaigncreateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignOperation is CampaignOperation Setter
// 返回数据
func (r *AlibabascbpadcampaigncreateAPIRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDto) error {
	r._campaignOperation = _campaignOperation
	r.Set("campaign_operation", _campaignOperation)
	return nil
}

// GetCampaignOperation CampaignOperation Getter
func (r AlibabascbpadcampaigncreateAPIRequest) GetCampaignOperation() *CampaignOperationDto {
	return r._campaignOperation
}
