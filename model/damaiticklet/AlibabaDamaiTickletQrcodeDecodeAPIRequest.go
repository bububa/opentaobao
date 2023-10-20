package damaiticklet

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaitickletqrcodedecodeAPIRequest 票夹-动态二维码-解码 API请求
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
type AlibabadamaitickletqrcodedecodeAPIRequest struct {
	model.Params
	// 加密二维码
	_encryptedQrCode string
	// 生产系统
	_productSystemId string
}

// NewAlibabadamaitickletqrcodedecodeRequest 初始化AlibabadamaitickletqrcodedecodeAPIRequest对象
func NewAlibabadamaitickletqrcodedecodeRequest() *AlibabadamaitickletqrcodedecodeAPIRequest {
	return &AlibabadamaitickletqrcodedecodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaitickletqrcodedecodeAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.ticklet.qrcode.decode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaitickletqrcodedecodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaitickletqrcodedecodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEncryptedQrCode is EncryptedQrCode Setter
// 加密二维码
func (r *AlibabadamaitickletqrcodedecodeAPIRequest) SetEncryptedQrCode(_encryptedQrCode string) error {
	r._encryptedQrCode = _encryptedQrCode
	r.Set("encrypted_qr_code", _encryptedQrCode)
	return nil
}

// GetEncryptedQrCode EncryptedQrCode Getter
func (r AlibabadamaitickletqrcodedecodeAPIRequest) GetEncryptedQrCode() string {
	return r._encryptedQrCode
}

// SetProductSystemId is ProductSystemId Setter
// 生产系统
func (r *AlibabadamaitickletqrcodedecodeAPIRequest) SetProductSystemId(_productSystemId string) error {
	r._productSystemId = _productSystemId
	r.Set("product_system_id", _productSystemId)
	return nil
}

// GetProductSystemId ProductSystemId Getter
func (r AlibabadamaitickletqrcodedecodeAPIRequest) GetProductSystemId() string {
	return r._productSystemId
}
