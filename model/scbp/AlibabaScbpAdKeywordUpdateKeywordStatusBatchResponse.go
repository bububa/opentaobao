package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改关键词状态 APIResponse
alibaba.scbp.ad.keyword.update.keyword.status.batch

修改关键词状态
*/
type AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_update_keyword_status_batch_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回错误集合
    
    ResultList   []ErrorKeyword `json:"result_list,omitempty" xml:"