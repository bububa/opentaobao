package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家地址库 API请求
taobao.logistics.address.search

通过此接口查询卖家地址库，
*/
type TaobaoLogisticsAddressSearchRequest struct {
    model.Params
    // 可选，参数列表如下：<br><font color='red'>no_def:查询非默认地址<br>get_def:查询默认取货地址，也即对应卖家后台地址库中发货地址（send_def暂不起作用）<br>cancel_def:查询默认退货地址<br>缺省此参数时，查询所有当前用户的地址库</font>
    _rdef   string
}

// 初始化TaobaoLogisticsAddressSearchRequest对象
func NewTaobaoLogisticsAddressSearchRequest() *TaobaoLogisticsAddressSearchRequest{
    return &TaobaoLogisticsAddressSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressSearchRequest) GetApiMethodName() string {
    return "taobao.logistics.address.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rdef Setter
// 可选，参数列表如下：<br><font color='red'>no_def:查询非默认地址<br>get_def:查询默认取货地址，也即对应卖家后台地址库中发货地址（send_def暂不起作用）<br>cancel_def:查询默认退货地址<br>缺省此参数时，查询所有当前用户的地址库</font>
func (r *TaobaoLogisticsAddressSearchRequest) SetRdef(_rdef string) error {
    r._rdef = _rdef
    r.Set("rdef", _rdef)
    return nil
}

// Rdef Getter
func (r TaobaoLogisticsAddressSearchRequest) GetRdef() string {
    return r._rdef
}
