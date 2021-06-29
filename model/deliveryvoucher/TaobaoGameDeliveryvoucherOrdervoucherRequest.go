package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预约接口 APIRequest
taobao.game.deliveryvoucher.ordervoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherOrdervoucherRequest struct {
    model.Params

    // 发券参数
    param0   *OrderVoucherRequest 

}

func NewTaobaoGameDeliveryvoucherOrdervoucherRequest() *TaobaoGameDeliveryvoucherOrdervoucherRequest{
    return &TaobaoGameDeliveryvoucherOrdervoucherRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherOrdervoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.ordervoucher"
}

func (r TaobaoGameDeliveryvoucherOrdervoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherOrdervoucherRequest) SetParam0(param0 *OrderVoucherRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherOrdervoucherRequest) GetParam0() *OrderVoucherRequest {
    return r.param0
}

