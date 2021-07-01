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
type TaobaoWlbImportsOrderCancelAPIRequest struct {
    model.Params
    // 物流订单编号
    _lgorderCode   string
}

// 初始化TaobaoWlbImportsOrderCancelAPIRequest对象
func NewTaobaoWlbImportsOrderCancelRequest() *TaobaoWlbImportsOrderCancelAPIRequest{
    return &TaobaoWlbImportsOrderCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LgorderCode Setter
// 物流订单编号
func (r *TaobaoWlbImportsOrderCancelAPIRequest) SetLgorderCode(_lgorderCode string) error {
    r._lgorderCode = _lgorderCode
    r.Set("lgorder_code", _lgorderCode)
    return nil
}

// LgorderCode Getter
func (r TaobaoWlbImportsOrderCancelAPIRequest) GetLgorderCode() string {
    return r._lgorderCode
}
