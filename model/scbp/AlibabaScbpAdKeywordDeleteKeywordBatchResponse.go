package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 APIResponse
alibaba.scbp.ad.keyword.delete.keyword.batch

删除关键词
*/
type AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_delete_keyword_batch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   int64 `json:"result,omitempty" xml:"