package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果查询 APIResponse
alibaba.alicom.order.preauthorize.query.fulfillment

预授权-履约结果查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomOrderPreauthorizeQueryFulfillmentResponse
}

type AlibabaAlicomOrderPreauthorizeQueryFulfillmentResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_query_fulfillment_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
