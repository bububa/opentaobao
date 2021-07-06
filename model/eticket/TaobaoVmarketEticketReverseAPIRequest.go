package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketReverseAPIRequest 电子凭证冲正接口 API请求
// taobao.vmarket.eticket.reverse
//
// 电子凭证平台冲正接口
type TaobaoVmarketEticketReverseAPIRequest struct {
	model.Params
	// 冲正的码，只支持单个码
	_reverseCode string
	// 需要冲正的核销记录对应核销流水号（对应的核销操作时候传递的自定义流水号）
	_consumeSecialNum string
	// 所有冲正后需要重新生成的码和对应的次数。码和次数之间用英文冒号分隔，多个码之间用英文逗号分隔。如果冲正后不需要重新生成码，留空
	_verifyCodes string
	// 不需要上传二维码图片或者冲正后不需要变更码的请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
	_qrImages string
	// 安全验证token，需要和该订单发码通知中的token一致
	_token string
	// 机具id，如果是码商必须传，如果是信任卖家不要传
	_posid string
	// 进行验码的电子凭证订单的订单ID
	_orderId int64
	// 冲正份数（必须是和被冲正的核销记录的份数一致）
	_reverseNum int64
	// 码商ID，是码商的话必须传递，如果是信任卖家不要传
	_codemerchantId int64
}

// NewTaobaoVmarketEticketReverseRequest 初始化TaobaoVmarketEticketReverseAPIRequest对象
func NewTaobaoVmarketEticketReverseRequest() *TaobaoVmarketEticketReverseAPIRequest {
	return &TaobaoVmarketEticketReverseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketReverseAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketReverseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReverseCode is ReverseCode Setter
// 冲正的码，只支持单个码
func (r *TaobaoVmarketEticketReverseAPIRequest) SetReverseCode(_reverseCode string) error {
	r._reverseCode = _reverseCode
	r.Set("reverse_code", _reverseCode)
	return nil
}

// GetReverseCode ReverseCode Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetReverseCode() string {
	return r._reverseCode
}

// SetConsumeSecialNum is ConsumeSecialNum Setter
// 需要冲正的核销记录对应核销流水号（对应的核销操作时候传递的自定义流水号）
func (r *TaobaoVmarketEticketReverseAPIRequest) SetConsumeSecialNum(_consumeSecialNum string) error {
	r._consumeSecialNum = _consumeSecialNum
	r.Set("consume_secial_num", _consumeSecialNum)
	return nil
}

// GetConsumeSecialNum ConsumeSecialNum Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetConsumeSecialNum() string {
	return r._consumeSecialNum
}

// SetVerifyCodes is VerifyCodes Setter
// 所有冲正后需要重新生成的码和对应的次数。码和次数之间用英文冒号分隔，多个码之间用英文逗号分隔。如果冲正后不需要重新生成码，留空
func (r *TaobaoVmarketEticketReverseAPIRequest) SetVerifyCodes(_verifyCodes string) error {
	r._verifyCodes = _verifyCodes
	r.Set("verify_codes", _verifyCodes)
	return nil
}

// GetVerifyCodes VerifyCodes Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetVerifyCodes() string {
	return r._verifyCodes
}

// SetQrImages is QrImages Setter
// 不需要上传二维码图片或者冲正后不需要变更码的请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaoVmarketEticketReverseAPIRequest) SetQrImages(_qrImages string) error {
	r._qrImages = _qrImages
	r.Set("qr_images", _qrImages)
	return nil
}

// GetQrImages QrImages Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetQrImages() string {
	return r._qrImages
}

// SetToken is Token Setter
// 安全验证token，需要和该订单发码通知中的token一致
func (r *TaobaoVmarketEticketReverseAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetToken() string {
	return r._token
}

// SetPosid is Posid Setter
// 机具id，如果是码商必须传，如果是信任卖家不要传
func (r *TaobaoVmarketEticketReverseAPIRequest) SetPosid(_posid string) error {
	r._posid = _posid
	r.Set("posid", _posid)
	return nil
}

// GetPosid Posid Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetPosid() string {
	return r._posid
}

// SetOrderId is OrderId Setter
// 进行验码的电子凭证订单的订单ID
func (r *TaobaoVmarketEticketReverseAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetReverseNum is ReverseNum Setter
// 冲正份数（必须是和被冲正的核销记录的份数一致）
func (r *TaobaoVmarketEticketReverseAPIRequest) SetReverseNum(_reverseNum int64) error {
	r._reverseNum = _reverseNum
	r.Set("reverse_num", _reverseNum)
	return nil
}

// GetReverseNum ReverseNum Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetReverseNum() int64 {
	return r._reverseNum
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID，是码商的话必须传递，如果是信任卖家不要传
func (r *TaobaoVmarketEticketReverseAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaoVmarketEticketReverseAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
