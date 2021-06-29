package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方入驻-开通服务 APIResponse
alibaba.alihealth.medical.item.publish

三方入驻-开通服务
*/
type AlibabaAlihealthMedicalItemPublishAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalItemPublishResponse
}

type AlibabaAlihealthMedicalItemPublishResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_item_publish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统返回的通用结果类
    
    Result1   *ServiceResult `json:"result1,omitempty" xml:"result1,omitempty"`

    
}
