package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户级别关键词推广状态 API返回值 
alibaba.scbp.account.status.get

查询账户级别关键词推广状态
*/
type AlibabaScbpAccountStatusGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAccountStatusGetResponse
}

// 查询账户级别关键词推广状态 成功返回结果
type AlibabaScbpAccountStatusGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_account_status_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // true:推广中,false:暂停
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
