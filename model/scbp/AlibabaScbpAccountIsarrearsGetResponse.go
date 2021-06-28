package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询关键词推广账户是否欠款 APIResponse
alibaba.scbp.account.isarrears.get

查询关键词推广账户是否欠款
*/
type AlibabaScbpAccountIsarrearsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_account_isarrears_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 客户的关键词推广账户是否欠款
    
    Result   bool `json:"result,omitempty" xml:"