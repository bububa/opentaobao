package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagEstimateUvAPIRequest 标签人群预估 API请求
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
type AlibabaScbpAdTargetTagEstimateUvAPIRequest struct {
	model.Params
	// 请求参数
	_tagEstimateOperation *TagEstimateOperationDto
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdTargetTagEstimateUvRequest 初始化AlibabaScbpAdTargetTagEstimateUvAPIRequest对象
func NewAlibabaScbpAdTargetTagEstimateUvRequest() *AlibabaScbpAdTargetTagEstimateUvAPIRequest {
	return &AlibabaScbpAdTargetTagEstimateUvAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.estimate.uv"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TagEstimateOperation Setter
// 请求参数
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) SetTagEstimateOperation(_tagEstimateOperation *TagEstimateOperationDto) error {
	r._tagEstimateOperation = _tagEstimateOperation
	r.Set("tag_estimate_operation", _tagEstimateOperation)
	return nil
}

// Get TagEstimateOperation Getter
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetTagEstimateOperation() *TagEstimateOperationDto {
	return r._tagEstimateOperation
}

// Set is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// Get TopContext Getter
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
