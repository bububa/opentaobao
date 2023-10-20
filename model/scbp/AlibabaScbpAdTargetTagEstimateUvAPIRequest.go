package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagEstimateUvAPIRequest 标签人群预估 API请求
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
type AlibabaScbpAdTargetTagEstimateUvAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_tagEstimateOperation *TagEstimateOperationDto
}

// NewAlibabaScbpAdTargetTagEstimateUvRequest 初始化AlibabaScbpAdTargetTagEstimateUvAPIRequest对象
func NewAlibabaScbpAdTargetTagEstimateUvRequest() *AlibabaScbpAdTargetTagEstimateUvAPIRequest {
	return &AlibabaScbpAdTargetTagEstimateUvAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) Reset() {
	r._topContext = nil
	r._tagEstimateOperation = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.target.tag.estimate.uv"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetTagEstimateOperation is TagEstimateOperation Setter
// 请求参数
func (r *AlibabaScbpAdTargetTagEstimateUvAPIRequest) SetTagEstimateOperation(_tagEstimateOperation *TagEstimateOperationDto) error {
	r._tagEstimateOperation = _tagEstimateOperation
	r.Set("tag_estimate_operation", _tagEstimateOperation)
	return nil
}

// GetTagEstimateOperation TagEstimateOperation Getter
func (r AlibabaScbpAdTargetTagEstimateUvAPIRequest) GetTagEstimateOperation() *TagEstimateOperationDto {
	return r._tagEstimateOperation
}

var poolAlibabaScbpAdTargetTagEstimateUvAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdTargetTagEstimateUvRequest()
	},
}

// GetAlibabaScbpAdTargetTagEstimateUvRequest 从 sync.Pool 获取 AlibabaScbpAdTargetTagEstimateUvAPIRequest
func GetAlibabaScbpAdTargetTagEstimateUvAPIRequest() *AlibabaScbpAdTargetTagEstimateUvAPIRequest {
	return poolAlibabaScbpAdTargetTagEstimateUvAPIRequest.Get().(*AlibabaScbpAdTargetTagEstimateUvAPIRequest)
}

// ReleaseAlibabaScbpAdTargetTagEstimateUvAPIRequest 将 AlibabaScbpAdTargetTagEstimateUvAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdTargetTagEstimateUvAPIRequest(v *AlibabaScbpAdTargetTagEstimateUvAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdTargetTagEstimateUvAPIRequest.Put(v)
}
