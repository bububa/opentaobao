package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推送及更新订单确认号 APIRequest
taobao.xhotel.order.update.confirmcode

商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。
*/
type TaobaoXhotelOrderUpdateConfirmcodeRequest struct {
    model.Params

    // 系统入参
    param   *UpdateOrderConfirmCodeParam 

}

func NewTaobaoXhotelOrderUpdateConfirmcodeRequest() *TaobaoXhotelOrderUpdateConfirmcodeRequest{
    return &TaobaoXhotelOrderUpdateConfirmcodeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderUpdateConfirmcodeRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.update.confirmcode"
}

func (r TaobaoXhotelOrderUpdateConfirmcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderUpdateConfirmcodeRequest) SetParam(param *UpdateOrderConfirmCodeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoXhotelOrderUpdateConfirmcodeRequest) GetParam() *UpdateOrderConfirmCodeParam {
    return r.param
}

