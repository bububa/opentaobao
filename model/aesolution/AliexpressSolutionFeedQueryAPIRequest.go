package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionfeedqueryAPIRequest aliexpress.solution.feed.query API请求
// aliexpress.solution.feed.query
//
// API for query the execution result of feed.
type AliexpresssolutionfeedqueryAPIRequest struct {
	model.Params
	// job id
	_jobId int64
}

// NewAliexpresssolutionfeedqueryRequest 初始化AliexpresssolutionfeedqueryAPIRequest对象
func NewAliexpresssolutionfeedqueryRequest() *AliexpresssolutionfeedqueryAPIRequest {
	return &AliexpresssolutionfeedqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionfeedqueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionfeedqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionfeedqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJobId is JobId Setter
// job id
func (r *AliexpresssolutionfeedqueryAPIRequest) SetJobId(_jobId int64) error {
	r._jobId = _jobId
	r.Set("job_id", _jobId)
	return nil
}

// GetJobId JobId Getter
func (r AliexpresssolutionfeedqueryAPIRequest) GetJobId() int64 {
	return r._jobId
}
