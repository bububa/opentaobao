package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketConsumeAPIRequest 电子票券消费通知 API请求
// taobao.vmarket.eticket.consume
//
// 外部合作商家电子票券消费回调接口
type TaobaoVmarketEticketConsumeAPIRequest struct {
	model.Params
	// 进行验码的电子凭证订单的订单ID
	_orderId int64
	// 核销的码，只支持单个码，多个码核销需要多次调用
	_verifyCode string
	// 核销份数
	_consumeNum int64
	// 安全验证token,需要和发码通知中的token一致
	_token string
	// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
	_codemerchantId int64
	// 机具ID(此参数信任卖家可不传递，码商必须传递)
	_posid string
	// 手机后四位(没有特殊说明请不要传该参数)
	_mobile string
	// 核销后需要重新生成的码，如果不需要重新生成码，不要传该参数
	_newCode string
	// 自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号
	_serialNum string
	// 不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。
	_qrImages string
}

// NewTaobaoVmarketEticketConsumeRequest 初始化TaobaoVmarketEticketConsumeAPIRequest对象
func NewTaobaoVmarketEticketConsumeRequest() *TaobaoVmarketEticketConsumeAPIRequest {
	return &TaobaoVmarketEticketConsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketConsumeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 进行验码的电子凭证订单的订单ID
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is VerifyCode Setter
// 核销的码，只支持单个码，多个码核销需要多次调用
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetVerifyCode(_verifyCode string) error {
	r._verifyCode = _verifyCode
	r.Set("verify_code", _verifyCode)
	return nil
}

// Get VerifyCode Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetVerifyCode() string {
	return r._verifyCode
}

// Set is ConsumeNum Setter
// 核销份数
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetConsumeNum(_consumeNum int64) error {
	r._consumeNum = _consumeNum
	r.Set("consume_num", _consumeNum)
	return nil
}

// Get ConsumeNum Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetConsumeNum() int64 {
	return r._consumeNum
}

// Set is Token Setter
// 安全验证token,需要和发码通知中的token一致
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetToken() string {
	return r._token
}

// Set is CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// Get CodemerchantId Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}

// Set is Posid Setter
// 机具ID(此参数信任卖家可不传递，码商必须传递)
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetPosid(_posid string) error {
	r._posid = _posid
	r.Set("posid", _posid)
	return nil
}

// Get Posid Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetPosid() string {
	return r._posid
}

// Set is Mobile Setter
// 手机后四位(没有特殊说明请不要传该参数)
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is NewCode Setter
// 核销后需要重新生成的码，如果不需要重新生成码，不要传该参数
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetNewCode(_newCode string) error {
	r._newCode = _newCode
	r.Set("new_code", _newCode)
	return nil
}

// Get NewCode Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetNewCode() string {
	return r._newCode
}

// Set is SerialNum Setter
// 自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetSerialNum(_serialNum string) error {
	r._serialNum = _serialNum
	r.Set("serial_num", _serialNum)
	return nil
}

// Get SerialNum Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetSerialNum() string {
	return r._serialNum
}

// Set is QrImages Setter
// 不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。
func (r *TaobaoVmarketEticketConsumeAPIRequest) SetQrImages(_qrImages string) error {
	r._qrImages = _qrImages
	r.Set("qr_images", _qrImages)
	return nil
}

// Get QrImages Getter
func (r TaobaoVmarketEticketConsumeAPIRequest) GetQrImages() string {
	return r._qrImages
}
