package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
记录用户每日消耗卡路里总量 APIResponse
alibaba.alihealth.alipaypfm.consume.record

记录用户每日消耗卡路里总量
*/
type AlibabaAlihealthAlipaypfmConsumeRecordAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthAlipaypfmConsumeRecordResponse
}

type AlibabaAlihealthAlipaypfmConsumeRecordResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_alipaypfm_consume_record_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和三方交互最外层model对象
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
