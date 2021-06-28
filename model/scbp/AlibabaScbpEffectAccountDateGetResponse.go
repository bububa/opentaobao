package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取最近报表生成时间 APIResponse
alibaba.scbp.effect.account.date.get

获取最近报表生成时间,格式为yyyy-MM-dd
*/
type AlibabaScbpEffectAccountDateGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpEffectAccountDateGetResponse
}

type AlibabaScbpEffectAccountDateGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_account_date_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最近生成报表的时间(US)
    
    ReportDate   string `json:"report_date,omitempty" xml:"report_date,omitempty"`

    
}
