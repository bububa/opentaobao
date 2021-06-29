package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝订单编号查询物流宝订单概要信息 API请求
taobao.wlb.wlborder.get

根据物流宝订单编号查询物流宝订单概要信息
*/
type TaobaoWlbWlborderGetRequest struct {
    model.Params
    // 物流宝订单编码
    wlbOrderCode   string
}

// 初始化TaobaoWlbWlborderGetRequest对象
func NewTaobaoWlbWlborderGetRequest() *TaobaoWlbWlborderGetRequest{
    return &TaobaoWlbWlborderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWlborderGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wlborder.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWlborderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WlbOrderCode Setter
// 物流宝订单编码
func (r *TaobaoWlbWlborderGetRequest) SetWlbOrderCode(wlbOrderCode string) error {
    r.wlbOrderCode = wlbOrderCode
    r.Set("wlb_order_code", wlbOrderCode)
    return nil
}

// WlbOrderCode Getter
func (r TaobaoWlbWlborderGetRequest) GetWlbOrderCode() string {
    return r.wlbOrderCode
}
