package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.query API请求
aliexpress.solution.feed.query

API for query the execution result of feed.
*/
type AliexpressSolutionFeedQueryRequest struct {
    model.Params
    // job id
    jobId   int64
}

// 初始化AliexpressSolutionFeedQueryRequest对象
func NewAliexpressSolutionFeedQueryRequest() *AliexpressSolutionFeedQueryRequest{
    return &AliexpressSolutionFeedQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedQueryRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// JobId Setter
// job id
func (r *AliexpressSolutionFeedQueryRequest) SetJobId(jobId int64) error {
    r.jobId = jobId
    r.Set("job_id", jobId)
    return nil
}

// JobId Getter
func (r AliexpressSolutionFeedQueryRequest) GetJobId() int64 {
    return r.jobId
}
