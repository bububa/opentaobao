package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
履约结果查询 APIResponse
alibaba.alicom.order.preauthorize.query.fulfillment

预授权-履约结果查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlicomOrderPreauthorizeQueryFulfillmentResponse `json:"alibaba_alicom_order_preauthorize_query_fulfillment_response,omitempty"`
}

type AlibabaAlicomOrderPreauthorizeQueryFulfillmentResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
