package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedInvalidateAPIRequest aliexpress.solution.feed.invalidate API请求
// aliexpress.solution.feed.invalidate
//
// Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
type AliexpressSolutionFeedInvalidateAPIRequest struct {
	model.Params
	// job id separated by ;  No more than 100 job Ids could be passed in one request.
	_jobIdList string
}

// NewAliexpressSolutionFeedInvalidateRequest 初始化AliexpressSolutionFeedInvalidateAPIRequest对象
func NewAliexpressSolutionFeedInvalidateRequest() *AliexpressSolutionFeedInvalidateAPIRequest {
	return &AliexpressSolutionFeedInvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedInvalidateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.invalidate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedInvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionFeedInvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJobIdList is JobIdList Setter
// job id separated by ;  No more than 100 job Ids could be passed in one request.
func (r *AliexpressSolutionFeedInvalidateAPIRequest) SetJobIdList(_jobIdList string) error {
	r._jobIdList = _jobIdList
	r.Set("job_id_list", _jobIdList)
	return nil
}

// GetJobIdList JobIdList Getter
func (r AliexpressSolutionFeedInvalidateAPIRequest) GetJobIdList() string {
	return r._jobIdList
}
