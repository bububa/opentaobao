package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卡券评价回传 APIRequest
taobao.game.deliveryvoucher.evaluate

卡券ISV回传商品评价
*/
type TaobaoGameDeliveryvoucherEvaluateRequest struct {
    model.Params

    // 系统自动生成
    param0   *VoucherEvaluateRequest 

}

func NewTaobaoGameDeliveryvoucherEvaluateRequest() *TaobaoGameDeliveryvoucherEvaluateRequest{
    return &TaobaoGameDeliveryvoucherEvaluateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherEvaluateRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.evaluate"
}

func (r TaobaoGameDeliveryvoucherEvaluateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherEvaluateRequest) SetParam0(param0 *VoucherEvaluateRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherEvaluateRequest) GetParam0() *VoucherEvaluateRequest {
    return r.param0
}

