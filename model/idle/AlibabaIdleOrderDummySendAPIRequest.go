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
type AlibabaIdleOrderDummySendAPIRequest struct {
    model.Params
    // 请求
    _paramOrderDummySendRequest   *OrderDummySendRequest
}

// 初始化AlibabaIdleOrderDummySendAPIRequest对象
func NewAlibabaIdleOrderDummySendRequest() *AlibabaIdleOrderDummySendAPIRequest{
    return &AlibabaIdleOrderDummySendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleOrderDummySendAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.order.dummy.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleOrderDummySendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderDummySendRequest Setter
// 请求
func (r *AlibabaIdleOrderDummySendAPIRequest) SetParamOrderDummySendRequest(_paramOrderDummySendRequest *OrderDummySendRequest) error {
    r._paramOrderDummySendRequest = _paramOrderDummySendRequest
    r.Set("param_order_dummy_send_request", _paramOrderDummySendRequest)
    return nil
}

// ParamOrderDummySendRequest Getter
func (r AlibabaIdleOrderDummySendAPIRequest) GetParamOrderDummySendRequest() *OrderDummySendRequest {
    return r._paramOrderDummySendRequest
}
