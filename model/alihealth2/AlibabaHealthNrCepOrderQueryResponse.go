package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查询接口 APIResponse
alibaba.health.nr.cep.order.query

订单详情查询接口
*/
type AlibabaHealthNrCepOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaHealthNrCepOrderQueryResponse
}

type AlibabaHealthNrCepOrderQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_health_nr_cep_order_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    ResponseResult   *ResponseResult `json:"response_result,omitempty" xml:"response_result,omitempty"`

    
}
