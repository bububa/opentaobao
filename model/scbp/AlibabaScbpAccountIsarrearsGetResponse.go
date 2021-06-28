package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词推广账户是否欠款 APIResponse
alibaba.scbp.account.isarrears.get

查询关键词推广账户是否欠款
*/
type AlibabaScbpAccountIsarrearsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAccountIsarrearsGetResponse `json:"alibaba_scbp_account_isarrears_get_response,omitempty"` 
    AlibabaScbpAccountIsarrearsGetResponse
}

/* model for simplify = false
type AlibabaScbpAccountIsarrearsGetResponse struct {

    // 客户的关键词推广账户是否欠款
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAccountIsarrearsGetResponse struct {

    // 客户的关键词推广账户是否欠款
    Result   bool `json:"result,omitempty"`

}
