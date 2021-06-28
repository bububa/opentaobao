package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车查询关键词 APIResponse
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词
*/
type AlibabaScbpAdKeywordGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"