package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的权益 API返回值 
alibaba.icbu.rfq.myequity

查询供应商权益接口
*/
type AlibabaIcbuRfqMyequityAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuRfqMyequityAPIResponseModel
}

// 我的权益 成功返回结果
type AlibabaIcbuRfqMyequityAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_icbu_rfq_myequity_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求返回结果
    ServiceResult   *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
