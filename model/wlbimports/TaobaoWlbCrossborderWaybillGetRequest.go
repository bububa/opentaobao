package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
集货商家pdf和云打印面单获取，pdf需要配置白名单 API请求
taobao.wlb.crossborder.waybill.get

【TOF】欧洲供应商PDF格式电子面单渲染下发
 需求链接：https://aone.alibaba-inc.com/req/21210808
*/
type TaobaoWlbCrossborderWaybillGetRequest struct {
    model.Params
    // 菜鸟物流单号
    _orderCode   string
}

// 初始化TaobaoWlbCrossborderWaybillGetRequest对象
func NewTaobaoWlbCrossborderWaybillGetRequest() *TaobaoWlbCrossborderWaybillGetRequest{
    return &TaobaoWlbCrossborderWaybillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbCrossborderWaybillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.crossborder.waybill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbCrossborderWaybillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 菜鸟物流单号
func (r *TaobaoWlbCrossborderWaybillGetRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbCrossborderWaybillGetRequest) GetOrderCode() string {
    return r._orderCode
}
