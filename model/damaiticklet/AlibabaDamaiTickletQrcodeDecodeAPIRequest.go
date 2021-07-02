package damaiticklet

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiTickletQrcodeDecodeAPIRequest 票夹-动态二维码-解码 API请求
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
type AlibabaDamaiTickletQrcodeDecodeAPIRequest struct {
	model.Params
	// 生产系统
	_productSystemId string
	// 加密二维码
	_encryptedQrCode string
}

// NewAlibabaDamaiTickletQrcodeDecodeRequest 初始化AlibabaDamaiTickletQrcodeDecodeAPIRequest对象
func NewAlibabaDamaiTickletQrcodeDecodeRequest() *AlibabaDamaiTickletQrcodeDecodeAPIRequest {
	return &AlibabaDamaiTickletQrcodeDecodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.ticklet.qrcode.decode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductSystemId Setter
// 生产系统
func (r *AlibabaDamaiTickletQrcodeDecodeAPIRequest) SetProductSystemId(_productSystemId string) error {
	r._productSystemId = _productSystemId
	r.Set("product_system_id", _productSystemId)
	return nil
}

// Get ProductSystemId Getter
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetProductSystemId() string {
	return r._productSystemId
}

// Set is EncryptedQrCode Setter
// 加密二维码
func (r *AlibabaDamaiTickletQrcodeDecodeAPIRequest) SetEncryptedQrCode(_encryptedQrCode string) error {
	r._encryptedQrCode = _encryptedQrCode
	r.Set("encrypted_qr_code", _encryptedQrCode)
	return nil
}

// Get EncryptedQrCode Getter
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetEncryptedQrCode() string {
	return r._encryptedQrCode
}
