package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
集货商家pdf和云打印面单获取，pdf需要配置白名单 APIRequest
taobao.wlb.crossborder.waybill.get

【TOF】欧洲供应商PDF格式电子面单渲染下发
 需求链接：https://aone.alibaba-inc.com/req/21210808
*/
type TaobaoWlbCrossborderWaybillGetRequest struct {
    model.Params

    // 菜鸟物流单号
    orderCode   string 

}

func NewTaobaoWlbCrossborderWaybillGetRequest() *TaobaoWlbCrossborderWaybillGetRequest{
    return &TaobaoWlbCrossborderWaybillGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbCrossborderWaybillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.crossborder.waybill.get"
}

func (r TaobaoWlbCrossborderWaybillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbCrossborderWaybillGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbCrossborderWaybillGetRequest) GetOrderCode() string {
    return r.orderCode
}

