package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车批量删除关键词 APIResponse
alibaba.scbp.ad.keyword.batchdelete

外贸直通车批量删除关键词
*/
type AlibabaScbpAdKeywordBatchdeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_batchdelete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除关键词是否成功
    
    Result   bool `json:"result,omitempty" xml:"