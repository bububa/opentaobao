package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运单信息 API请求
taobao.wlb.imports.waybill.get

一般进口商家，获取订单的电子面单链接地址
*/
type TaobaoWlbImportsWaybillGetRequest struct {
    model.Params
    // 物流订单号
    _orderCode   string
}

// 初始化TaobaoWlbImportsWaybillGetRequest对象
func NewTaobaoWlbImportsWaybillGetRequest() *TaobaoWlbImportsWaybillGetRequest{
    return &TaobaoWlbImportsWaybillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsWaybillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.waybill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsWaybillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 物流订单号
func (r *TaobaoWlbImportsWaybillGetRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbImportsWaybillGetRequest) GetOrderCode() string {
    return r._orderCode
}
