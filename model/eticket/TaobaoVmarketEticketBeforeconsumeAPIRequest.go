package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketbeforeconsumeAPIRequest 电子凭证验码前置确认 API请求
// taobao.vmarket.eticket.beforeconsume
//
// 商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作
type TaobaovmarketeticketbeforeconsumeAPIRequest struct {
	model.Params
	// 需要验的码
	_verifyCode string
	// 安全验证token，需要和发码通知中的token一致
	_token string
	// 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
	_posid string
	// 手机号码后四位,没有特殊说明请不要传
	_mobile string
	// 需要验码的电子凭证订单ID
	_orderId int64
	// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
	_codemerchantId int64
}

// NewTaobaovmarketeticketbeforeconsumeRequest 初始化TaobaovmarketeticketbeforeconsumeAPIRequest对象
func NewTaobaovmarketeticketbeforeconsumeRequest() *TaobaovmarketeticketbeforeconsumeAPIRequest {
	return &TaobaovmarketeticketbeforeconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.beforeconsume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyCode is VerifyCode Setter
// 需要验的码
func (r *TaobaovmarketeticketbeforeconsumeAPIRequest) SetVerifyCode(_verifyCode string) error {
	r._verifyCode = _verifyCode
	r.Set("verify_code", _verifyCode)
	return nil
}

// GetVerifyCode VerifyCode Getter
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetVerifyCode() string {
	return r._verifyCode
}

// SetToken is Token Setter
// 安全验证token，需要和发码通知中的token一致
func (r *TaobaovmarketeticketbeforeconsumeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetToken() string {
	return r._token
}

// SetPosid is Posid Setter
// 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
func (r *TaobaovmarketeticketbeforeconsumeAPIRequest) SetPosid(_posid string) error {
	r._posid = _posid
	r.Set("posid", _posid)
	return nil
}

// GetPosid Posid Getter
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetPosid() string {
	return r._posid
}

// SetMobile is Mobile Setter
// 手机号码后四位,没有特殊说明请不要传
func (r *TaobaovmarketeticketbeforeconsumeAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetMobile() string {
	return r._mobile
}

// SetOrderId is OrderId Setter
// 需要验码的电子凭证订单ID
func (r *TaobaovmarketeticketbeforeconsumeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
func (r *TaobaovmarketeticketbeforeconsumeAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaovmarketeticketbeforeconsumeAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
