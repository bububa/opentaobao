package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词添加 APIResponse
alibaba.scbp.ad.keyword.create.keyword.batch

关键词添加
*/
type AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_create_keyword_batch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回错误集合
    
    ResultList   []ErrorKeyword `json:"result_list,omitempty" xml:"