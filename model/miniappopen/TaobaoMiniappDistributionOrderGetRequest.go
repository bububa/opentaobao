package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序投放-查询小程序投放计划信息 API请求
taobao.miniapp.distribution.order.get

服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
*/
type TaobaoMiniappDistributionOrderGetRequest struct {
    model.Params
    // 查询入参
    orderIdRequest   *DistributionOrderQueryByIdOpenRequest
}

// 初始化TaobaoMiniappDistributionOrderGetRequest对象
func NewTaobaoMiniappDistributionOrderGetRequest() *TaobaoMiniappDistributionOrderGetRequest{
    return &TaobaoMiniappDistributionOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMiniappDistributionOrderGetRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMiniappDistributionOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderIdRequest Setter
// 查询入参
func (r *TaobaoMiniappDistributionOrderGetRequest) SetOrderIdRequest(orderIdRequest *DistributionOrderQueryByIdOpenRequest) error {
    r.orderIdRequest = orderIdRequest
    r.Set("order_id_request", orderIdRequest)
    return nil
}

// OrderIdRequest Getter
func (r TaobaoMiniappDistributionOrderGetRequest) GetOrderIdRequest() *DistributionOrderQueryByIdOpenRequest {
    return r.orderIdRequest
}
