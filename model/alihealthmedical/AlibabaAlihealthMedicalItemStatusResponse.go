package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品上下架 APIResponse
alibaba.alihealth.medical.item.status

医生通三方机构平台进行服务商品上下架操作
*/
type AlibabaAlihealthMedicalItemStatusAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalItemStatusResponse
}

type AlibabaAlihealthMedicalItemStatusResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_item_status_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Shelfresult   *ServiceResult `json:"shelfresult,omitempty" xml:"shelfresult,omitempty"`

    
}
