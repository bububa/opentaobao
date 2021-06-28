package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
履约结果 APIResponse
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果
*/
type AlibabaAlicomOrderPreauthorizeFulfillmentAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alicom_order_preauthorize_fulfillment_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"