package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消物流宝订单 API请求
taobao.wlb.order.cancel

取消物流宝订单
*/
type TaobaoWlbOrderCancelAPIRequest struct {
    model.Params
    // 物流宝订单编号
    _wlbOrderCode   string
}

// 初始化TaobaoWlbOrderCancelAPIRequest对象
func NewTaobaoWlbOrderCancelRequest() *TaobaoWlbOrderCancelAPIRequest{
    return &TaobaoWlbOrderCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderCancelAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WlbOrderCode Setter
// 物流宝订单编号
func (r *TaobaoWlbOrderCancelAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
    r._wlbOrderCode = _wlbOrderCode
    r.Set("wlb_order_code", _wlbOrderCode)
    return nil
}

// WlbOrderCode Getter
func (r TaobaoWlbOrderCancelAPIRequest) GetWlbOrderCode() string {
    return r._wlbOrderCode
}
