package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方机构查询订单详情接口 APIResponse
alibaba.alihealth.medical.order.query

查询订单详情，包括评价
*/
type AlibabaAlihealthMedicalOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalOrderQueryResponse
}

type AlibabaAlihealthMedicalOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
