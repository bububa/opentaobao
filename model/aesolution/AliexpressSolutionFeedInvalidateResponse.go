package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.invalidate APIResponse
aliexpress.solution.feed.invalidate

Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
*/
type AliexpressSolutionFeedInvalidateAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionFeedInvalidateResponse
}

type AliexpressSolutionFeedInvalidateResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_feed_invalidate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // successful list
    
    SuccessfulList   []int64 `json:"successful_list,omitempty" xml:"successful_list>int64,omitempty"`
    
    
    // no need to be invalidated jobs. Usually caused by that the job is being or has already been executed.
    
    NoNeedInvalidationList   []int64 `json:"no_need_invalidation_list,omitempty" xml:"no_need_invalidation_list>int64,omitempty"`
    
    
    // failed list
    
    FailedList   []int64 `json:"failed_list,omitempty" xml:"failed_list>int64,omitempty"`
    
    
}
