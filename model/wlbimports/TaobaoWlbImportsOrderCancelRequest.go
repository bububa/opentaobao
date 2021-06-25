package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一般进口取消物流订单 APIRequest
taobao.wlb.imports.order.cancel

商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
*/
type TaobaoWlbImportsOrderCancelRequest struct {
    model.Params

    // 物流订单编号
    lgorderCode   string 

}

func NewTaobaoWlbImportsOrderCancelRequest() *TaobaoWlbImportsOrderCancelRequest{
    return &TaobaoWlbImportsOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportsOrderCancelRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.order.cancel"
}

func (r TaobaoWlbImportsOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportsOrderCancelRequest) SetLgorderCode(lgorderCode string) error {
    r.lgorderCode = lgorderCode
    r.Set("lgorder_code", lgorderCode)
    return nil
}

func (r TaobaoWlbImportsOrderCancelRequest) GetLgorderCode() string {
    return r.lgorderCode
}

