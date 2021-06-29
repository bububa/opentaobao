package alihealth

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康处方平台获取授权码 APIResponse
alibaba.alihealth.prescription.auth.get

获取处方用户授权
*/
type AlibabaAlihealthPrescriptionAuthGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthPrescriptionAuthGetResponse
}

type AlibabaAlihealthPrescriptionAuthGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_prescription_auth_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
