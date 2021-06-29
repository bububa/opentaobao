package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 API请求
taobao.qimen.order.cancel

ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
*/
type TaobaoQimenOrderCancelRequest struct {
    model.Params
    // 
    _request   *OrderCancelRequest
}

// 初始化TaobaoQimenOrderCancelRequest对象
func NewTaobaoQimenOrderCancelRequest() *TaobaoQimenOrderCancelRequest{
    return &TaobaoQimenOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderCancelRequest) GetApiMethodName() string {
    return "taobao.qimen.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderCancelRequest) SetRequest(_request *OrderCancelRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderCancelRequest) GetRequest() *OrderCancelRequest {
    return r._request
}
