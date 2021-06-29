package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单履约V1 API请求
alibaba.idle.recycle.order.fulfillment

外部回收商针对自有回收订单的履行
*/
type AlibabaIdleRecycleOrderFulfillmentRequest struct {
    model.Params
    // 订单同步入参
    _param0   *RecycleOrderSynDto
}

// 初始化AlibabaIdleRecycleOrderFulfillmentRequest对象
func NewAlibabaIdleRecycleOrderFulfillmentRequest() *AlibabaIdleRecycleOrderFulfillmentRequest{
    return &AlibabaIdleRecycleOrderFulfillmentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.fulfillment"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 订单同步入参
func (r *AlibabaIdleRecycleOrderFulfillmentRequest) SetParam0(_param0 *RecycleOrderSynDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaIdleRecycleOrderFulfillmentRequest) GetParam0() *RecycleOrderSynDto {
    return r._param0
}
