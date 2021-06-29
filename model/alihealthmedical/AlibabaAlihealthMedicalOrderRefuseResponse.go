package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构通知平台"医生拒诊" APIResponse
alibaba.alihealth.medical.order.refuse

三方机构通知平台"医生拒诊"
*/
type AlibabaAlihealthMedicalOrderRefuseAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalOrderRefuseResponse
}

type AlibabaAlihealthMedicalOrderRefuseResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_order_refuse_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
