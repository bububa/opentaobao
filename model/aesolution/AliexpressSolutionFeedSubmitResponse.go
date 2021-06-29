package aesolution

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.submit APIResponse
aliexpress.solution.feed.submit

API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
*/
type AliexpressSolutionFeedSubmitAPIResponse struct {
    model.CommonResponse
    AliexpressSolutionFeedSubmitResponse
}

type AliexpressSolutionFeedSubmitResponse struct {
    XMLName xml.Name `xml:"aliexpress_solution_feed_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // job id,which is for querying the job response later.
    
    JobId   int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`

    
}
