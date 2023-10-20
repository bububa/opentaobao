package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest 删除屏蔽词 API请求
// alibaba.scbp.ad.campaign.delete.forbidden.keyword
//
// 删除屏蔽词
type AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
	// 请求参数
	_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto
}

// NewAlibabascbpadcampaigndeleteforbiddenkeywordRequest 初始化AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest对象
func NewAlibabascbpadcampaigndeleteforbiddenkeywordRequest() *AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest {
	return &AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.campaign.delete.forbidden.keyword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetForbiddenKeywordBatchOperation is ForbiddenKeywordBatchOperation Setter
// 请求参数
func (r *AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) SetForbiddenKeywordBatchOperation(_forbiddenKeywordBatchOperation *ForbiddenKeywordBatchOperationDto) error {
	r._forbiddenKeywordBatchOperation = _forbiddenKeywordBatchOperation
	r.Set("forbidden_keyword_batch_operation", _forbiddenKeywordBatchOperation)
	return nil
}

// GetForbiddenKeywordBatchOperation ForbiddenKeywordBatchOperation Getter
func (r AlibabascbpadcampaigndeleteforbiddenkeywordAPIRequest) GetForbiddenKeywordBatchOperation() *ForbiddenKeywordBatchOperationDto {
	return r._forbiddenKeywordBatchOperation
}
