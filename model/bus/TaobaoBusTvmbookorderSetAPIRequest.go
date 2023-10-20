package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmbookordersetAPIRequest 线下自助机通知出票接口 API请求
// taobao.bus.tvmbookorder.set
//
// 出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
type TaobaobustvmbookordersetAPIRequest struct {
	model.Params
	// 乘客票面信息
	_passengers []TvmPassengerVo
	// 飞猪订单号
	_alitripOrderId string
	// 出票时间 2017-03-03 11:22:33
	_bookTime string
	// 取票人手机号
	_fetchPhone string
	// 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
	_payMode string
	// 检票口
	_ticketGate string
	// true代表出票成功；false代表出票失败
	_success bool
	// 是否支持电子票
	_supportEticket bool
}

// NewTaobaobustvmbookordersetRequest 初始化TaobaobustvmbookordersetAPIRequest对象
func NewTaobaobustvmbookordersetRequest() *TaobaobustvmbookordersetAPIRequest {
	return &TaobaobustvmbookordersetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmbookordersetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmbookorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmbookordersetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmbookordersetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPassengers is Passengers Setter
// 乘客票面信息
func (r *TaobaobustvmbookordersetAPIRequest) SetPassengers(_passengers []TvmPassengerVo) error {
	r._passengers = _passengers
	r.Set("passengers", _passengers)
	return nil
}

// GetPassengers Passengers Getter
func (r TaobaobustvmbookordersetAPIRequest) GetPassengers() []TvmPassengerVo {
	return r._passengers
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaobustvmbookordersetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobustvmbookordersetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

// SetBookTime is BookTime Setter
// 出票时间 2017-03-03 11:22:33
func (r *TaobaobustvmbookordersetAPIRequest) SetBookTime(_bookTime string) error {
	r._bookTime = _bookTime
	r.Set("book_time", _bookTime)
	return nil
}

// GetBookTime BookTime Getter
func (r TaobaobustvmbookordersetAPIRequest) GetBookTime() string {
	return r._bookTime
}

// SetFetchPhone is FetchPhone Setter
// 取票人手机号
func (r *TaobaobustvmbookordersetAPIRequest) SetFetchPhone(_fetchPhone string) error {
	r._fetchPhone = _fetchPhone
	r.Set("fetch_phone", _fetchPhone)
	return nil
}

// GetFetchPhone FetchPhone Getter
func (r TaobaobustvmbookordersetAPIRequest) GetFetchPhone() string {
	return r._fetchPhone
}

// SetPayMode is PayMode Setter
// 取值范围  ALIPAY （飞猪渠道）; WECHAT（微信）; BANKCARD（银行卡）;CASH（现金）; OWN_ALIPAY（自身支付宝渠道，非飞猪渠道）
func (r *TaobaobustvmbookordersetAPIRequest) SetPayMode(_payMode string) error {
	r._payMode = _payMode
	r.Set("pay_mode", _payMode)
	return nil
}

// GetPayMode PayMode Getter
func (r TaobaobustvmbookordersetAPIRequest) GetPayMode() string {
	return r._payMode
}

// SetTicketGate is TicketGate Setter
// 检票口
func (r *TaobaobustvmbookordersetAPIRequest) SetTicketGate(_ticketGate string) error {
	r._ticketGate = _ticketGate
	r.Set("ticket_gate", _ticketGate)
	return nil
}

// GetTicketGate TicketGate Getter
func (r TaobaobustvmbookordersetAPIRequest) GetTicketGate() string {
	return r._ticketGate
}

// SetSuccess is Success Setter
// true代表出票成功；false代表出票失败
func (r *TaobaobustvmbookordersetAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r TaobaobustvmbookordersetAPIRequest) GetSuccess() bool {
	return r._success
}

// SetSupportEticket is SupportEticket Setter
// 是否支持电子票
func (r *TaobaobustvmbookordersetAPIRequest) SetSupportEticket(_supportEticket bool) error {
	r._supportEticket = _supportEticket
	r.Set("support_eticket", _supportEticket)
	return nil
}

// GetSupportEticket SupportEticket Getter
func (r TaobaobustvmbookordersetAPIRequest) GetSupportEticket() bool {
	return r._supportEticket
}
