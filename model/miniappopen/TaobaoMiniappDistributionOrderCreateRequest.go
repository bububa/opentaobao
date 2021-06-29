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
type TaobaoMiniappDistributionOrderCreateRequest struct {
    model.Params
    // 投放计划信息
    orderRequest   *DistributionOrderSaveOpenRequest
}

// 初始化TaobaoMiniappDistributionOrderCreateRequest对象
func NewTaobaoMiniappDistributionOrderCreateRequest() *TaobaoMiniappDistributionOrderCreateRequest{
    return &TaobaoMiniappDistributionOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderCreateRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderRequest Setter
// 投放计划信息
func (r *TaobaoMiniappDistributionOrderCreateRequest) SetOrderRequest(orderRequest *DistributionOrderSaveOpenRequest) error {
    r.orderRequest = orderRequest
    r.Set("order_request", orderRequest)
    return nil
}

// OrderRequest Getter
func (r TaobaoMiniappDistributionOrderCreateRequest) GetOrderRequest() *DistributionOrderSaveOpenRequest {
    return r.orderRequest
}
