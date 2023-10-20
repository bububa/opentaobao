package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedQueryAPIRequest aliexpress.solution.feed.query API请求
// aliexpress.solution.feed.query
//
// API for query the execution result of feed.
type AliexpressSolutionFeedQueryAPIRequest struct {
	model.Params
	// job id
	_jobId int64
}

// NewAliexpressSolutionFeedQueryRequest 初始化AliexpressSolutionFeedQueryAPIRequest对象
func NewAliexpressSolutionFeedQueryRequest() *AliexpressSolutionFeedQueryAPIRequest {
	return &AliexpressSolutionFeedQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionFeedQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJobId is JobId Setter
// job id
func (r *AliexpressSolutionFeedQueryAPIRequest) SetJobId(_jobId int64) error {
	r._jobId = _jobId
	r.Set("job_id", _jobId)
	return nil
}

// GetJobId JobId Getter
func (r AliexpressSolutionFeedQueryAPIRequest) GetJobId() int64 {
	return r._jobId
}
