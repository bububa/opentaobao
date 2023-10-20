package eticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketResendAPIRequest 外部合作商家重发电子凭证回调接口 API请求
// taobao.vmarket.eticket.resend
//
// 外部合作商家重发电子凭证回调接口
type TaobaoVmarketEticketResendAPIRequest struct {
	model.Params
	// 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
	_verifyCodes string
	// 安全验证token,回传淘宝发通知时发过来的token串
	_token string
	// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
	_qrImages string
	// 订单编号
	_orderId int64
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
}

// NewTaobaoVmarketEticketResendRequest 初始化TaobaoVmarketEticketResendAPIRequest对象
func NewTaobaoVmarketEticketResendRequest() *TaobaoVmarketEticketResendAPIRequest {
	return &TaobaoVmarketEticketResendAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketResendAPIRequest) Reset() {
	r._verifyCodes = ""
	r._token = ""
	r._qrImages = ""
	r._orderId = 0
	r._codemerchantId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketResendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketResendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketResendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyCodes is VerifyCodes Setter
// 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
func (r *TaobaoVmarketEticketResendAPIRequest) SetVerifyCodes(_verifyCodes string) error {
	r._verifyCodes = _verifyCodes
	r.Set("verify_codes", _verifyCodes)
	return nil
}

// GetVerifyCodes VerifyCodes Getter
func (r TaobaoVmarketEticketResendAPIRequest) GetVerifyCodes() string {
	return r._verifyCodes
}

// SetToken is Token Setter
// 安全验证token,回传淘宝发通知时发过来的token串
func (r *TaobaoVmarketEticketResendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoVmarketEticketResendAPIRequest) GetToken() string {
	return r._token
}

// SetQrImages is QrImages Setter
// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaoVmarketEticketResendAPIRequest) SetQrImages(_qrImages string) error {
	r._qrImages = _qrImages
	r.Set("qr_images", _qrImages)
	return nil
}

// GetQrImages QrImages Getter
func (r TaobaoVmarketEticketResendAPIRequest) GetQrImages() string {
	return r._qrImages
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *TaobaoVmarketEticketResendAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketResendAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaoVmarketEticketResendAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaoVmarketEticketResendAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}

var poolTaobaoVmarketEticketResendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketResendRequest()
	},
}

// GetTaobaoVmarketEticketResendRequest 从 sync.Pool 获取 TaobaoVmarketEticketResendAPIRequest
func GetTaobaoVmarketEticketResendAPIRequest() *TaobaoVmarketEticketResendAPIRequest {
	return poolTaobaoVmarketEticketResendAPIRequest.Get().(*TaobaoVmarketEticketResendAPIRequest)
}

// ReleaseTaobaoVmarketEticketResendAPIRequest 将 TaobaoVmarketEticketResendAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketResendAPIRequest(v *TaobaoVmarketEticketResendAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketResendAPIRequest.Put(v)
}
