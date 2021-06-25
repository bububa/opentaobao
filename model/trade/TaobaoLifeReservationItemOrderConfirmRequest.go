package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生服购后预约单外部确认 APIRequest
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
*/
type TaobaoLifeReservationItemOrderConfirmRequest struct {
    model.Params

    // 淘宝主单号
    tradeNo   string 

    // 凭证ID
    ticketId   string 

    // 审核类型，PASS-通过；REJECT-驳回
    optType   string 

}

func NewTaobaoLifeReservationItemOrderConfirmRequest() *TaobaoLifeReservationItemOrderConfirmRequest{
    return &TaobaoLifeReservationItemOrderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLifeReservationItemOrderConfirmRequest) GetApiMethodName() string {
    return "taobao.life.reservation.item.order.confirm"
}

func (r TaobaoLifeReservationItemOrderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLifeReservationItemOrderConfirmRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r TaobaoLifeReservationItemOrderConfirmRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *TaobaoLifeReservationItemOrderConfirmRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

func (r TaobaoLifeReservationItemOrderConfirmRequest) GetTicketId() string {
    return r.ticketId
}

func (r *TaobaoLifeReservationItemOrderConfirmRequest) SetOptType(optType string) error {
    r.optType = optType
    r.Set("opt_type", optType)
    return nil
}

func (r TaobaoLifeReservationItemOrderConfirmRequest) GetOptType() string {
    return r.optType
}

