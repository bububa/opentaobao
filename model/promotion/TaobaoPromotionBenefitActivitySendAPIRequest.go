package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionBenefitActivitySendAPIRequest 活动权益发放接口 API请求
// taobao.promotion.benefit.activity.send
//
// 活动权益发放接口，用于卖家针对活动进行权益发放
type TaobaoPromotionBenefitActivitySendAPIRequest struct {
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

// NewTaobaoPromotionBenefitActivitySendRequest 初始化TaobaoPromotionBenefitActivitySendAPIRequest对象
func NewTaobaoPromotionBenefitActivitySendRequest() *TaobaoPromotionBenefitActivitySendAPIRequest {
	return &TaobaoPromotionBenefitActivitySendAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) Reset() {
	r._nick = ""
	r._platNick = ""
	r._mixReceiverId = ""
	r._ouid = ""
	r._uid = ""
	r._sendRequest = nil
	r._receiverId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.benefit.activity.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 混淆的接收者nick
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetNick() string {
	return r._nick
}

// SetPlatNick is PlatNick Setter
// 非混淆的接收者nick
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetPlatNick(_platNick string) error {
	r._platNick = _platNick
	r.Set("plat_nick", _platNick)
	return nil
}

// GetPlatNick PlatNick Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetPlatNick() string {
	return r._platNick
}

// SetMixReceiverId is MixReceiverId Setter
// 混淆的接收者id
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetMixReceiverId(_mixReceiverId string) error {
	r._mixReceiverId = _mixReceiverId
	r.Set("mix_receiver_id", _mixReceiverId)
	return nil
}

// GetMixReceiverId MixReceiverId Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetMixReceiverId() string {
	return r._mixReceiverId
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetOuid() string {
	return r._ouid
}

// SetUid is Uid Setter
// openid
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetUid() string {
	return r._uid
}

// SetSendRequest is SendRequest Setter
// 单个权益发放请求
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetSendRequest(_sendRequest *BenefitSingleSendRequest) error {
	r._sendRequest = _sendRequest
	r.Set("send_request", _sendRequest)
	return nil
}

// GetSendRequest SendRequest Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetSendRequest() *BenefitSingleSendRequest {
	return r._sendRequest
}

// SetReceiverId is ReceiverId Setter
// 非混淆的接收者id
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetReceiverId(_receiverId int64) error {
	r._receiverId = _receiverId
	r.Set("receiver_id", _receiverId)
	return nil
}

// GetReceiverId ReceiverId Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetReceiverId() int64 {
	return r._receiverId
}

var poolTaobaoPromotionBenefitActivitySendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionBenefitActivitySendRequest()
	},
}

// GetTaobaoPromotionBenefitActivitySendRequest 从 sync.Pool 获取 TaobaoPromotionBenefitActivitySendAPIRequest
func GetTaobaoPromotionBenefitActivitySendAPIRequest() *TaobaoPromotionBenefitActivitySendAPIRequest {
	return poolTaobaoPromotionBenefitActivitySendAPIRequest.Get().(*TaobaoPromotionBenefitActivitySendAPIRequest)
}

// ReleaseTaobaoPromotionBenefitActivitySendAPIRequest 将 TaobaoPromotionBenefitActivitySendAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionBenefitActivitySendAPIRequest(v *TaobaoPromotionBenefitActivitySendAPIRequest) {
	v.Reset()
	poolTaobaoPromotionBenefitActivitySendAPIRequest.Put(v)
}
