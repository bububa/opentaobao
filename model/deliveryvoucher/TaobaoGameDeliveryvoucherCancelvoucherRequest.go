package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
作废券 APIRequest
taobao.game.deliveryvoucher.cancelvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherCancelvoucherRequest struct {
    model.Params

    // 发券参数
    param0   *CancelVoucherRequest 

}

func NewTaobaoGameDeliveryvoucherCancelvoucherRequest() *TaobaoGameDeliveryvoucherCancelvoucherRequest{
    return &TaobaoGameDeliveryvoucherCancelvoucherRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherCancelvoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.cancelvoucher"
}

func (r TaobaoGameDeliveryvoucherCancelvoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherCancelvoucherRequest) SetParam0(param0 *CancelVoucherRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherCancelvoucherRequest) GetParam0() *CancelVoucherRequest {
    return r.param0
}

