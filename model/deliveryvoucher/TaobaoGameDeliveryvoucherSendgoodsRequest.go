package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提货券发货接口 APIRequest
taobao.game.deliveryvoucher.sendgoods

提货券发券接口：同步券和订单的关联信息
*/
type TaobaoGameDeliveryvoucherSendgoodsRequest struct {
    model.Params

    // 发券参数
    param0   *SendVoucherRequest 

}

func NewTaobaoGameDeliveryvoucherSendgoodsRequest() *TaobaoGameDeliveryvoucherSendgoodsRequest{
    return &TaobaoGameDeliveryvoucherSendgoodsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherSendgoodsRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.sendgoods"
}

func (r TaobaoGameDeliveryvoucherSendgoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherSendgoodsRequest) SetParam0(param0 *SendVoucherRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherSendgoodsRequest) GetParam0() *SendVoucherRequest {
    return r.param0
}

