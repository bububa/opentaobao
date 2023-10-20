package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketresendAPIRequest 外部合作商家重发电子凭证回调接口 API请求
// taobao.vmarket.eticket.resend
//
// 外部合作商家重发电子凭证回调接口
type TaobaovmarketeticketresendAPIRequest struct {
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

// NewTaobaovmarketeticketresendRequest 初始化TaobaovmarketeticketresendAPIRequest对象
func NewTaobaovmarketeticketresendRequest() *TaobaovmarketeticketresendAPIRequest {
	return &TaobaovmarketeticketresendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketresendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.resend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketresendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketresendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVerifyCodes is VerifyCodes Setter
// 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
func (r *TaobaovmarketeticketresendAPIRequest) SetVerifyCodes(_verifyCodes string) error {
	r._verifyCodes = _verifyCodes
	r.Set("verify_codes", _verifyCodes)
	return nil
}

// GetVerifyCodes VerifyCodes Getter
func (r TaobaovmarketeticketresendAPIRequest) GetVerifyCodes() string {
	return r._verifyCodes
}

// SetToken is Token Setter
// 安全验证token,回传淘宝发通知时发过来的token串
func (r *TaobaovmarketeticketresendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaovmarketeticketresendAPIRequest) GetToken() string {
	return r._token
}

// SetQrImages is QrImages Setter
// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaovmarketeticketresendAPIRequest) SetQrImages(_qrImages string) error {
	r._qrImages = _qrImages
	r.Set("qr_images", _qrImages)
	return nil
}

// GetQrImages QrImages Getter
func (r TaobaovmarketeticketresendAPIRequest) GetQrImages() string {
	return r._qrImages
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *TaobaovmarketeticketresendAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketeticketresendAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaovmarketeticketresendAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaovmarketeticketresendAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
