package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传入出库单api APIResponse
alibaba.alihealth.tracecodeseller.bill.upload

上传入出库单api
*/
type AlibabaAlihealthTracecodesellerBillUploadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthTracecodesellerBillUploadResponse
}

type AlibabaAlihealthTracecodesellerBillUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_bill_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // model
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
