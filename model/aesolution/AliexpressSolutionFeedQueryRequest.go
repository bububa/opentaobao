package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.query APIRequest
aliexpress.solution.feed.query

API for query the execution result of feed.
*/
type AliexpressSolutionFeedQueryRequest struct {
    model.Params

    // job id
    jobId   int64 

}

func NewAliexpressSolutionFeedQueryRequest() *AliexpressSolutionFeedQueryRequest{
    return &AliexpressSolutionFeedQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionFeedQueryRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.query"
}

func (r AliexpressSolutionFeedQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionFeedQueryRequest) SetJobId(jobId int64) error {
    r.jobId = jobId
    r.Set("job_id", jobId)
    return nil
}

func (r AliexpressSolutionFeedQueryRequest) GetJobId() int64 {
    return r.jobId
}

