package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无需物流发货 API请求
alibaba.idle.order.dummy.send

适用于电子卡券等虚拟商品不需要物流的商品发货。
*/
type AlibabaIdleOrderDummySendRequest struct {
    model.Params
    // 请求
    paramOrderDummySendRequest   *OrderDummySendRequest
}

// 初始化AlibabaIdleOrderDummySendRequest对象
func NewAlibabaIdleOrderDummySendRequest() *AlibabaIdleOrderDummySendRequest{
    return &AlibabaIdleOrderDummySendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleOrderDummySendRequest) GetApiMethodName() string {
    return "alibaba.idle.order.dummy.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleOrderDummySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderDummySendRequest Setter
// 请求
func (r *AlibabaIdleOrderDummySendRequest) SetParamOrderDummySendRequest(paramOrderDummySendRequest *OrderDummySendRequest) error {
    r.paramOrderDummySendRequest = paramOrderDummySendRequest
    r.Set("param_order_dummy_send_request", paramOrderDummySendRequest)
    return nil
}

// ParamOrderDummySendRequest Getter
func (r AlibabaIdleOrderDummySendRequest) GetParamOrderDummySendRequest() *OrderDummySendRequest {
    return r.paramOrderDummySendRequest
}
