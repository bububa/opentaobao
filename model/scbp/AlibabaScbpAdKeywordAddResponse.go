package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车加词 APIResponse
alibaba.scbp.ad.keyword.add

外贸直通车加词服务
*/
type AlibabaScbpAdKeywordAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求加入的词
    
    Keyword   string `json:"keyword,omitempty" xml:"