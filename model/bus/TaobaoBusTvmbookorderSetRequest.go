package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机通知出票接口 API请求
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

// 初始化TaobaoBusTvmbookorderSetRequest对象
func NewTaobaoBusTvmbookorderSetRequest() *TaobaoBusTvmbookorderSetRequest{
    return &TaobaoBusTvmbookorderSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmbookorderSetRequest) GetApiMethodName() string {
    return "taobao.bus.tvmbookorder.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmbookorderSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmbookorderSetRequest) SetAlitripOrderId(alitripOrderId string) error {
    r.alitripOrderId = alitripOrderId
    r.Set("alitrip_order_id", alitripOrderId)
    return nil
}

// AlitripOrderId Getter
func (r TaobaoBusTvmbookorderSetRequest) GetAlitripOrderId() string {
    return r.alitripOrderId
}
// BookTime Setter
// 出票时间 2017-03-03 11:22:33
func (r *TaobaoBusTvmbookorderSetRequest) SetBookTime(bookTime string) error {
    r.bookTime = bookTime
    r.Set("book_time", bookTime)
    return nil
}

// BookTime Getter
func (r TaobaoBusTvmbookorderSetRequest) GetBookTime() string {
    return r.bookTime
}
// Success Setter
// true代表出票成功；false代表出票失败
func (r *TaobaoBusTvmbookorderSetRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

// Success Getter
func (r TaobaoBusTvmbookorderSetRequest) GetSuccess() bool {
    return r.success
}
// PayMode Setter
// 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
func (r *TaobaoBusTvmbookorderSetRequest) SetPayMode(payMode string) error {
    r.payMode = payMode
    r.Set("pay_mode", payMode)
    return nil
}

// PayMode Getter
func (r TaobaoBusTvmbookorderSetRequest) GetPayMode() string {
    return r.payMode
}
// FetchPhone Setter
// 取票人手机号
func (r *TaobaoBusTvmbookorderSetRequest) SetFetchPhone(fetchPhone string) error {
    r.fetchPhone = fetchPhone
    r.Set("fetch_phone", fetchPhone)
    return nil
}

// FetchPhone Getter
func (r TaobaoBusTvmbookorderSetRequest) GetFetchPhone() string {
    return r.fetchPhone
}
// Passengers Setter
// 乘客票面信息
func (r *TaobaoBusTvmbookorderSetRequest) SetPassengers(passengers []TvmPassengerVo) error {
    r.passengers = passengers
    r.Set("passengers", passengers)
    return nil
}

// Passengers Getter
func (r TaobaoBusTvmbookorderSetRequest) GetPassengers() []TvmPassengerVo {
    return r.passengers
}
// SupportEticket Setter
// 是否支持电子票
func (r *TaobaoBusTvmbookorderSetRequest) SetSupportEticket(supportEticket bool) error {
    r.supportEticket = supportEticket
    r.Set("support_eticket", supportEticket)
    return nil
}

// SupportEticket Getter
func (r TaobaoBusTvmbookorderSetRequest) GetSupportEticket() bool {
    return r.supportEticket
}
// TicketGate Setter
// 检票口
func (r *TaobaoBusTvmbookorderSetRequest) SetTicketGate(ticketGate string) error {
    r.ticketGate = ticketGate
    r.Set("ticket_gate", ticketGate)
    return nil
}

// TicketGate Getter
func (r TaobaoBusTvmbookorderSetRequest) GetTicketGate() string {
    return r.ticketGate
}
