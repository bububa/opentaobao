package deliveryvoucher

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
监控预约数据 APIRequest
taobao.game.deliveryvoucher.watch

监控预约数据
*/
type TaobaoGameDeliveryvoucherWatchRequest struct {
    model.Params

    // 入参
    param0   *WatchAppointmentRequest 

}

func NewTaobaoGameDeliveryvoucherWatchRequest() *TaobaoGameDeliveryvoucherWatchRequest{
    return &TaobaoGameDeliveryvoucherWatchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoGameDeliveryvoucherWatchRequest) GetApiMethodName() string {
    return "taobao.game.deliveryvoucher.watch"
}

func (r TaobaoGameDeliveryvoucherWatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoGameDeliveryvoucherWatchRequest) SetParam0(param0 *WatchAppointmentRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoGameDeliveryvoucherWatchRequest) GetParam0() *WatchAppointmentRequest {
    return r.param0
}

