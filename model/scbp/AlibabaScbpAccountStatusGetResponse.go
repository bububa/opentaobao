package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户级别关键词推广状态 APIResponse
alibaba.scbp.account.status.get

查询账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_account_status_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true:推广中,false:暂停
    
    Result   bool `json:"result,omitempty" xml:"