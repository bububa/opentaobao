package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivitysendAPIRequest 活动权益发放接口 API请求
// taobao.promotion.benefit.activity.send
//
// 活动权益发放接口，用于卖家针对活动进行权益发放
type TaobaopromotionbenefitactivitysendAPIRequest struct {
	model.Params
	// 混淆的接收者nick
	_nick string
	// 非混淆的接收者nick
	_platNick string
	// 混淆的接收者id
	_mixReceiverId string
	// ouid
	_ouid string
	// openid
	_uid string
	// 单个权益发放请求
	_sendRequest *BenefitSingleSendRequest
	// 非混淆的接收者id
	_receiverId int64
}

// NewTaobaopromotionbenefitactivitysendRequest 初始化TaobaopromotionbenefitactivitysendAPIRequest对象
func NewTaobaopromotionbenefitactivitysendRequest() *TaobaopromotionbenefitactivitysendAPIRequest {
	return &TaobaopromotionbenefitactivitysendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 混淆的接收者nick
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetNick() string {
	return r._nick
}

// SetPlatNick is PlatNick Setter
// 非混淆的接收者nick
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetPlatNick(_platNick string) error {
	r._platNick = _platNick
	r.Set("plat_nick", _platNick)
	return nil
}

// GetPlatNick PlatNick Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetPlatNick() string {
	return r._platNick
}

// SetMixReceiverId is MixReceiverId Setter
// 混淆的接收者id
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetMixReceiverId(_mixReceiverId string) error {
	r._mixReceiverId = _mixReceiverId
	r.Set("mix_receiver_id", _mixReceiverId)
	return nil
}

// GetMixReceiverId MixReceiverId Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetMixReceiverId() string {
	return r._mixReceiverId
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetOuid() string {
	return r._ouid
}

// SetUid is Uid Setter
// openid
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetUid() string {
	return r._uid
}

// SetSendRequest is SendRequest Setter
// 单个权益发放请求
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetSendRequest(_sendRequest *BenefitSingleSendRequest) error {
	r._sendRequest = _sendRequest
	r.Set("send_request", _sendRequest)
	return nil
}

// GetSendRequest SendRequest Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetSendRequest() *BenefitSingleSendRequest {
	return r._sendRequest
}

// SetReceiverId is ReceiverId Setter
// 非混淆的接收者id
func (r *TaobaopromotionbenefitactivitysendAPIRequest) SetReceiverId(_receiverId int64) error {
	r._receiverId = _receiverId
	r.Set("receiver_id", _receiverId)
	return nil
}

// GetReceiverId ReceiverId Getter
func (r TaobaopromotionbenefitactivitysendAPIRequest) GetReceiverId() int64 {
	return r._receiverId
}
