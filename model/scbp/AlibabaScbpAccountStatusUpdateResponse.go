package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改账户级别关键词推广状态 APIResponse
alibaba.scbp.account.status.update

修改账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAccountStatusUpdateResponse
}

type AlibabaScbpAccountStatusUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_account_status_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
