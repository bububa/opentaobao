package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询账户级别关键词推广状态 APIResponse
alibaba.scbp.account.status.get

查询账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAccountStatusGetResponse `json:"alibaba_scbp_account_status_get_response,omitempty"` 
    AlibabaScbpAccountStatusGetResponse
}

/* model for simplify = false
type AlibabaScbpAccountStatusGetResponse struct {

    // true:推广中,false:暂停
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAccountStatusGetResponse struct {

    // true:推广中,false:暂停
    Result   bool `json:"result,omitempty"`

}
