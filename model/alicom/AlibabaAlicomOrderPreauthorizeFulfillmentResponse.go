package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果 API返回值 
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果
*/
type AlibabaAlicomOrderPreauthorizeFulfillmentAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomOrderPreauthorizeFulfillmentResponse
}

// 履约结果 成功返回结果
type AlibabaAlicomOrderPreauthorizeFulfillmentResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_order_preauthorize_fulfillment_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
