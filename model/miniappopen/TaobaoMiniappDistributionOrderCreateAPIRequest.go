package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建小程序投放计划 API请求
taobao.miniapp.distribution.order.create

帮助商家，创建小程序的投放计划。
*/
type TaobaoMiniappDistributionOrderCreateAPIRequest struct {
    model.Params
    // 投放计划信息
    _orderRequest   *DistributionOrderSaveOpenRequest
}

// 初始化TaobaoMiniappDistributionOrderCreateAPIRequest对象
func NewTaobaoMiniappDistributionOrderCreateRequest() *TaobaoMiniappDistributionOrderCreateAPIRequest{
    return &TaobaoMiniappDistributionOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderRequest Setter
// 投放计划信息
func (r *TaobaoMiniappDistributionOrderCreateAPIRequest) SetOrderRequest(_orderRequest *DistributionOrderSaveOpenRequest) error {
    r._orderRequest = _orderRequest
    r.Set("order_request", _orderRequest)
    return nil
}

// OrderRequest Getter
func (r TaobaoMiniappDistributionOrderCreateAPIRequest) GetOrderRequest() *DistributionOrderSaveOpenRequest {
    return r._orderRequest
}
