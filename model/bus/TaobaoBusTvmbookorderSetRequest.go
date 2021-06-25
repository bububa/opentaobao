package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机通知出票接口 APIRequest
taobao.bus.tvmbookorder.set

出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
*/
type TaobaoBusTvmbookorderSetRequest struct {
    model.Params

    // 飞猪订单号
    alitripOrderId   string 

    // 出票时间 2017-03-03 11:22:33
    bookTime   string 

    // true代表出票成功；false代表出票失败
    success   bool 

    // 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
    payMode   string 

    // 取票人手机号
    fetchPhone   string 

    // 乘客票面信息
    passengers   []TvmPassengerVo 

    // 是否支持电子票
    supportEticket   bool 

    // 检票口
    ticketGate   string 

}

func NewTaobaoBusTvmbookorderSetRequest() *TaobaoBusTvmbookorderSetRequest{
    return &TaobaoBusTvmbookorderSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusTvmbookorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmbookorder.set"
}

func (r TaobaoBusTvmbookorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusTvmbookorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}

func (r *TaobaoBusTvmbookorderSetRequest) SetBookTime(bookTime string) error {
    r.bookTime = bookTime
    r.Set("book_time", bookTime)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetBookTime() string {
    return r.bookTime
}

func (r *TaobaoBusTvmbookorderSetRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetSuccess() bool {
    return r.success
}

func (r *TaobaoBusTvmbookorderSetRequest) SetPayMode(payMode string) error {
    r.payMode = payMode
    r.Set("pay_mode", payMode)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetPayMode() string {
    return r.payMode
}

func (r *TaobaoBusTvmbookorderSetRequest) SetFetchPhone(fetchPhone string) error {
    r.fetchPhone = fetchPhone
    r.Set("fetch_phone", fetchPhone)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetFetchPhone() string {
    return r.fetchPhone
}

func (r *TaobaoBusTvmbookorderSetRequest) SetPassengers(passengers []TvmPassengerVo) error {
    r.passengers = passengers
    r.Set("passengers", passengers)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetPassengers() []TvmPassengerVo {
    return r.passengers
}

func (r *TaobaoBusTvmbookorderSetRequest) SetSupportEticket(supportEticket bool) error {
    r.supportEticket = supportEticket
    r.Set("support_eticket", supportEticket)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetSupportEticket() bool {
    return r.supportEticket
}

func (r *TaobaoBusTvmbookorderSetRequest) SetTicketGate(ticketGate string) error {
    r.ticketGate = ticketGate
    r.Set("ticket_gate", ticketGate)
    return nil
}

func (r TaobaoBusTvmbookorderSetRequest) GetTicketGate() string {
    return r.ticketGate
}

