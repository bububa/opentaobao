package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketsendAPIRequest 商家电子凭证发码成功回调接口 API请求
// taobao.vmarket.eticket.send
//
// 外部商家成功发码回调接口
type TaobaovmarketeticketsendAPIRequest struct {
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

// NewTaobaovmarketeticketsendRequest 初始化TaobaovmarketeticketsendAPIRequest对象
func NewTaobaovmarketeticketsendRequest() *TaobaovmarketeticketsendAPIRequest {
	return &TaobaovmarketeticketsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketsendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyCodes is VerifyCodes Setter
// 发送成功的验证码及可验证次数的列表，码和可验证次数用英文冒号分隔，多个码之间用英文逗号分隔，所有字符都为英文半角
func (r *TaobaovmarketeticketsendAPIRequest) SetVerifyCodes(_verifyCodes string) error {
	r._verifyCodes = _verifyCodes
	r.Set("verify_codes", _verifyCodes)
	return nil
}

// GetVerifyCodes VerifyCodes Getter
func (r TaobaovmarketeticketsendAPIRequest) GetVerifyCodes() string {
	return r._verifyCodes
}

// SetToken is Token Setter
// 安全验证token，需要和发码通知中的token一致
func (r *TaobaovmarketeticketsendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaovmarketeticketsendAPIRequest) GetToken() string {
	return r._token
}

// SetQrImages is QrImages Setter
// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaovmarketeticketsendAPIRequest) SetQrImages(_qrImages string) error {
	r._qrImages = _qrImages
	r.Set("qr_images", _qrImages)
	return nil
}

// GetQrImages QrImages Getter
func (r TaobaovmarketeticketsendAPIRequest) GetQrImages() string {
	return r._qrImages
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *TaobaovmarketeticketsendAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketeticketsendAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家,不需要传
func (r *TaobaovmarketeticketsendAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaovmarketeticketsendAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
