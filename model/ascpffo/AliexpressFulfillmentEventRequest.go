package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE履约事件处理 API请求
aliexpress.fulfillment.event

AE用 履约底层声明发货能力
*/
type AliexpressFulfillmentEventRequest struct {
    model.Params
    // 入参对象
    _param   *FulfillmentOrderStatusUpdateRequest
}

// 初始化AliexpressFulfillmentEventRequest对象
func NewAliexpressFulfillmentEventRequest() *AliexpressFulfillmentEventRequest{
    return &AliexpressFulfillmentEventRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressFulfillmentEventRequest) GetApiMethodName() string {
    return "aliexpress.fulfillment.event"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressFulfillmentEventRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AliexpressFulfillmentEventRequest) SetParam(_param *FulfillmentOrderStatusUpdateRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AliexpressFulfillmentEventRequest) GetParam() *FulfillmentOrderStatusUpdateRequest {
    return r._param
}
