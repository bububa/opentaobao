package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
快易通零售B2C APIResponse
alibaba.alihealth.drug.kyt.uploadb2cbill

快易通零售B2C单据上传
*/
type AlibabaAlihealthDrugKytUploadb2cbillAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytUploadb2cbillResponse
}

type AlibabaAlihealthDrugKytUploadb2cbillResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadb2cbill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果值
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 调用结果
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 调用结果描述
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // success
    
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`

    
}
