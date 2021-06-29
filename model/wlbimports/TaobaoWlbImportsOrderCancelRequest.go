package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一般进口取消物流订单 API请求
taobao.wlb.imports.order.cancel

商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
*/
type TaobaoWlbImportsOrderCancelRequest struct {
    model.Params
    // 物流订单编号
    _lgorderCode   string
}

// 初始化TaobaoWlbImportsOrderCancelRequest对象
func NewTaobaoWlbImportsOrderCancelRequest() *TaobaoWlbImportsOrderCancelRequest{
    return &TaobaoWlbImportsOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsOrderCancelRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LgorderCode Setter
// 物流订单编号
func (r *TaobaoWlbImportsOrderCancelRequest) SetLgorderCode(_lgorderCode string) error {
    r._lgorderCode = _lgorderCode
    r.Set("lgorder_code", _lgorderCode)
    return nil
}

// LgorderCode Getter
func (r TaobaoWlbImportsOrderCancelRequest) GetLgorderCode() string {
    return r._lgorderCode
}
