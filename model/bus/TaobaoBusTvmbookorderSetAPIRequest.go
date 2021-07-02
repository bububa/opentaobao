package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmbookorderSetAPIRequest 线下自助机通知出票接口 API请求
// taobao.bus.tvmbookorder.set
//
// 出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
type TaobaoBusTvmbookorderSetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
	// 出票时间 2017-03-03 11:22:33
	_bookTime string
	// true代表出票成功；false代表出票失败
	_success bool
	// 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
	_payMode string
	// 取票人手机号
	_fetchPhone string
	// 乘客票面信息
	_passengers []TvmPassengerVo
	// 是否支持电子票
	_supportEticket bool
	// 检票口
	_ticketGate string
}

// NewTaobaoBusTvmbookorderSetRequest 初始化TaobaoBusTvmbookorderSetAPIRequest对象
func NewTaobaoBusTvmbookorderSetRequest() *TaobaoBusTvmbookorderSetAPIRequest {
	return &TaobaoBusTvmbookorderSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmbookorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmbookorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmbookorderSetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// Get AlitripOrderId Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

// Set is BookTime Setter
// 出票时间 2017-03-03 11:22:33
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetBookTime(_bookTime string) error {
	r._bookTime = _bookTime
	r.Set("book_time", _bookTime)
	return nil
}

// Get BookTime Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetBookTime() string {
	return r._bookTime
}

// Set is Success Setter
// true代表出票成功；false代表出票失败
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// Get Success Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetSuccess() bool {
	return r._success
}

// Set is PayMode Setter
// 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetPayMode(_payMode string) error {
	r._payMode = _payMode
	r.Set("pay_mode", _payMode)
	return nil
}

// Get PayMode Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetPayMode() string {
	return r._payMode
}

// Set is FetchPhone Setter
// 取票人手机号
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetFetchPhone(_fetchPhone string) error {
	r._fetchPhone = _fetchPhone
	r.Set("fetch_phone", _fetchPhone)
	return nil
}

// Get FetchPhone Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetFetchPhone() string {
	return r._fetchPhone
}

// Set is Passengers Setter
// 乘客票面信息
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetPassengers(_passengers []TvmPassengerVo) error {
	r._passengers = _passengers
	r.Set("passengers", _passengers)
	return nil
}

// Get Passengers Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetPassengers() []TvmPassengerVo {
	return r._passengers
}

// Set is SupportEticket Setter
// 是否支持电子票
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetSupportEticket(_supportEticket bool) error {
	r._supportEticket = _supportEticket
	r.Set("support_eticket", _supportEticket)
	return nil
}

// Get SupportEticket Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetSupportEticket() bool {
	return r._supportEticket
}

// Set is TicketGate Setter
// 检票口
func (r *TaobaoBusTvmbookorderSetAPIRequest) SetTicketGate(_ticketGate string) error {
	r._ticketGate = _ticketGate
	r.Set("ticket_gate", _ticketGate)
	return nil
}

// Get TicketGate Getter
func (r TaobaoBusTvmbookorderSetAPIRequest) GetTicketGate() string {
	return r._ticketGate
}
