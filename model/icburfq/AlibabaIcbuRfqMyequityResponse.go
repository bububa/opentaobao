package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的权益 APIResponse
alibaba.icbu.rfq.myequity

查询供应商权益接口
*/
type AlibabaIcbuRfqMyequityAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuRfqMyequityResponse
}

type AlibabaIcbuRfqMyequityResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_rfq_myequity_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求返回结果
    
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`

    
}
