package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口 APIResponse
alibaba.alicom.order.exchange.create

代理商调用该接口来进行积分兑换
*/
type AlibabaAlicomOrderExchangeCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlicomOrderExchangeCreateResponse `json:"alibaba_alicom_order_exchange_create_response,omitempty"`
}

type AlibabaAlicomOrderExchangeCreateResponse struct {

    // result
    Result   *CommonResult `json:"result,omitempty"`

}
