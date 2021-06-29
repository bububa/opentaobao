package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.invalidate API请求
aliexpress.solution.feed.invalidate

Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
*/
type AliexpressSolutionFeedInvalidateRequest struct {
    model.Params
    // job id separated by ;  No more than 100 job Ids could be passed in one request.
    _jobIdList   string
}

// 初始化AliexpressSolutionFeedInvalidateRequest对象
func NewAliexpressSolutionFeedInvalidateRequest() *AliexpressSolutionFeedInvalidateRequest{
    return &AliexpressSolutionFeedInvalidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedInvalidateRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.invalidate"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedInvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// JobIdList Setter
// job id separated by ;  No more than 100 job Ids could be passed in one request.
func (r *AliexpressSolutionFeedInvalidateRequest) SetJobIdList(_jobIdList string) error {
    r._jobIdList = _jobIdList
    r.Set("job_id_list", _jobIdList)
    return nil
}

// JobIdList Getter
func (r AliexpressSolutionFeedInvalidateRequest) GetJobIdList() string {
    return r._jobIdList
}
