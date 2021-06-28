package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改账户级别关键词推广状态 APIResponse
alibaba.scbp.account.status.update

修改账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAccountStatusUpdateResponse `json:"alibaba_scbp_account_status_update_response,omitempty"` 
    AlibabaScbpAccountStatusUpdateResponse
}

/* model for simplify = false
type AlibabaScbpAccountStatusUpdateResponse struct {

    // 修改成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAccountStatusUpdateResponse struct {

    // 修改成功
    Result   bool `json:"result,omitempty"`

}
