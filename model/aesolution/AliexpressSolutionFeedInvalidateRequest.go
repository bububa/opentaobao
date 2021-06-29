package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.invalidate APIRequest
aliexpress.solution.feed.invalidate

Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
*/
type AliexpressSolutionFeedInvalidateRequest struct {
    model.Params

    // job id separated by ;  No more than 100 job Ids could be passed in one request.
    jobIdList   string 

}

func NewAliexpressSolutionFeedInvalidateRequest() *AliexpressSolutionFeedInvalidateRequest{
    return &AliexpressSolutionFeedInvalidateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSolutionFeedInvalidateRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.invalidate"
}

func (r AliexpressSolutionFeedInvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSolutionFeedInvalidateRequest) SetJobIdList(jobIdList string) error {
    r.jobIdList = jobIdList
    r.Set("job_id_list", jobIdList)
    return nil
}

func (r AliexpressSolutionFeedInvalidateRequest) GetJobIdList() string {
    return r.jobIdList
}

