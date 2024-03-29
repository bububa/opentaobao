package eticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketSendAPIRequest 商家电子凭证发码成功回调接口 API请求
// taobao.vmarket.eticket.send
//
// 外部商家成功发码回调接口
type TaobaoVmarketEticketSendAPIRequest struct {
	model.Params
	// 发送成功的验证码及可验证次数的列表，码和可验证次数用英文冒号分隔，多个码之间用英文逗号分隔，所有字符都为英文半角
	_verifyCodes string
	// 安全验证token，需要和发码通知中的token一致
	_token string
	// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
	_qrImages string
	// 订单编号
	_orderId int64
	// 码商ID,是码商的话必须传递,如果是信任卖家,不需要传
	_codemerchantId int64
}

// NewTaobaoVmarketEticketSendRequest 初始化TaobaoVmarketEticketSendAPIRequest对象
func NewTaobaoVmarketEticketSendRequest() *TaobaoVmarketEticketSendAPIRequest {
	return &TaobaoVmarketEticketSendAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketSendAPIRequest) Reset() {
	r._verifyCodes = ""
	r._token = ""
	r._qrImages = ""
	r._orderId = 0
	r._codemerchantId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketSendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyCodes is VerifyCodes Setter
// 发送成功的验证码及可验证次数的列表，码和可验证次数用英文冒号分隔，多个码之间用英文逗号分隔，所有字符都为英文半角
func (r *TaobaoVmarketEticketSendAPIRequest) SetVerifyCodes(_verifyCodes string) error {
	r._verifyCodes = _verifyCodes
	r.Set("verify_codes", _verifyCodes)
	return nil
}

// GetVerifyCodes VerifyCodes Getter
func (r TaobaoVmarketEticketSendAPIRequest) GetVerifyCodes() string {
	return r._verifyCodes
}

// SetToken is Token Setter
// 安全验证token，需要和发码通知中的token一致
func (r *TaobaoVmarketEticketSendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoVmarketEticketSendAPIRequest) GetToken() string {
	return r._token
}

// SetQrImages is QrImages Setter
// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaoVmarketEticketSendAPIRequest) SetQrImages(_qrImages string) error {
	r._qrImages = _qrImages
	r.Set("qr_images", _qrImages)
	return nil
}

// GetQrImages QrImages Getter
func (r TaobaoVmarketEticketSendAPIRequest) GetQrImages() string {
	return r._qrImages
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *TaobaoVmarketEticketSendAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketSendAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家,不需要传
func (r *TaobaoVmarketEticketSendAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaoVmarketEticketSendAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}

var poolTaobaoVmarketEticketSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketSendRequest()
	},
}

// GetTaobaoVmarketEticketSendRequest 从 sync.Pool 获取 TaobaoVmarketEticketSendAPIRequest
func GetTaobaoVmarketEticketSendAPIRequest() *TaobaoVmarketEticketSendAPIRequest {
	return poolTaobaoVmarketEticketSendAPIRequest.Get().(*TaobaoVmarketEticketSendAPIRequest)
}

// ReleaseTaobaoVmarketEticketSendAPIRequest 将 TaobaoVmarketEticketSendAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketSendAPIRequest(v *TaobaoVmarketEticketSendAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketSendAPIRequest.Put(v)
}
