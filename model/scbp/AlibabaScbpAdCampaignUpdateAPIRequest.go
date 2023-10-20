package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaignupdateAPIRequest 修改计划 API请求
// alibaba.scbp.ad.campaign.update
//
// 修改计划
type AlibabascbpadcampaignupdateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 修改数据
	_campaignOperation *CampaignOperationDto
}

// NewAlibabascbpadcampaignupdateRequest 初始化AlibabascbpadcampaignupdateAPIRequest对象
func NewAlibabascbpadcampaignupdateRequest() *AlibabascbpadcampaignupdateAPIRequest {
	return &AlibabascbpadcampaignupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaignupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaignupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaignupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaignupdateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaignupdateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignOperation is CampaignOperation Setter
// 修改数据
func (r *AlibabascbpadcampaignupdateAPIRequest) SetCampaignOperation(_campaignOperation *CampaignOperationDto) error {
	r._campaignOperation = _campaignOperation
	r.Set("campaign_operation", _campaignOperation)
	return nil
}

// GetCampaignOperation CampaignOperation Getter
func (r AlibabascbpadcampaignupdateAPIRequest) GetCampaignOperation() *CampaignOperationDto {
	return r._campaignOperation
}
