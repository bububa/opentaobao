package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提货券发券接口 APIRequest
taobao.game.deliveryvoucher.sendvoucher

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherSendvoucherRequest struct {
    model.Params

    // 发券参数
    param0   *SendVoucherRequest 

}

func NewTaobaoGameDeliveryvoucherSendvoucherRequest() *TaobaoGameDeliveryvoucherSendvoucherRequest{
    return &TaobaoGameDeliveryvoucherSendvoucherRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherSendvoucherRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.sendvoucher"
}

func (r TaobaoGameDeliveryvoucherSendvoucherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherSendvoucherRequest) SetParam0(param0 *SendVoucherRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherSendvoucherRequest) GetParam0() *SendVoucherRequest {
    return r.param0
}

