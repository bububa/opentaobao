package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据挂起（恢复）接口 API请求
taobao.qimen.order.pending

ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
*/
type TaobaoQimenOrderPendingRequest struct {
    model.Params
    // 
    request   *OrderPendingRequest
}

// 初始化TaobaoQimenOrderPendingRequest对象
func NewTaobaoQimenOrderPendingRequest() *TaobaoQimenOrderPendingRequest{
    return &TaobaoQimenOrderPendingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderPendingRequest) GetApiMethodName() string {
    return "taobao.qimen.order.pending"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderPendingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderPendingRequest) SetRequest(request *OrderPendingRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderPendingRequest) GetRequest() *OrderPendingRequest {
    return r.request
}
