package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadtargettagestimateuvAPIRequest 标签人群预估 API请求
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
type AlibabascbpadtargettagestimateuvAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_tagEstimateOperation *TagEstimateOperationDto
}

// NewAlibabascbpadtargettagestimateuvRequest 初始化AlibabascbpadtargettagestimateuvAPIRequest对象
func NewAlibabascbpadtargettagestimateuvRequest() *AlibabascbpadtargettagestimateuvAPIRequest {
	return &AlibabascbpadtargettagestimateuvAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadtargettagestimateuvAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.estimate.uv"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadtargettagestimateuvAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadtargettagestimateuvAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadtargettagestimateuvAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadtargettagestimateuvAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetTagEstimateOperation is TagEstimateOperation Setter
// 请求参数
func (r *AlibabascbpadtargettagestimateuvAPIRequest) SetTagEstimateOperation(_tagEstimateOperation *TagEstimateOperationDto) error {
	r._tagEstimateOperation = _tagEstimateOperation
	r.Set("tag_estimate_operation", _tagEstimateOperation)
	return nil
}

// GetTagEstimateOperation TagEstimateOperation Getter
func (r AlibabascbpadtargettagestimateuvAPIRequest) GetTagEstimateOperation() *TagEstimateOperationDto {
	return r._tagEstimateOperation
}
