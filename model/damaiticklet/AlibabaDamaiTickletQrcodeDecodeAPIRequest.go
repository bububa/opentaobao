package damaiticklet

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiTickletQrcodeDecodeAPIRequest 票夹-动态二维码-解码 API请求
// alibaba.damai.ticklet.qrcode.decode
//
// 对于票夹的动态二维码进行解码
type AlibabaDamaiTickletQrcodeDecodeAPIRequest struct {
	model.Params
	// 加密二维码
	_encryptedQrCode string
	// 生产系统
	_productSystemId string
}

// NewAlibabaDamaiTickletQrcodeDecodeRequest 初始化AlibabaDamaiTickletQrcodeDecodeAPIRequest对象
func NewAlibabaDamaiTickletQrcodeDecodeRequest() *AlibabaDamaiTickletQrcodeDecodeAPIRequest {
	return &AlibabaDamaiTickletQrcodeDecodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiTickletQrcodeDecodeAPIRequest) Reset() {
	r._encryptedQrCode = ""
	r._productSystemId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.ticklet.qrcode.decode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEncryptedQrCode is EncryptedQrCode Setter
// 加密二维码
func (r *AlibabaDamaiTickletQrcodeDecodeAPIRequest) SetEncryptedQrCode(_encryptedQrCode string) error {
	r._encryptedQrCode = _encryptedQrCode
	r.Set("encrypted_qr_code", _encryptedQrCode)
	return nil
}

// GetEncryptedQrCode EncryptedQrCode Getter
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetEncryptedQrCode() string {
	return r._encryptedQrCode
}

// SetProductSystemId is ProductSystemId Setter
// 生产系统
func (r *AlibabaDamaiTickletQrcodeDecodeAPIRequest) SetProductSystemId(_productSystemId string) error {
	r._productSystemId = _productSystemId
	r.Set("product_system_id", _productSystemId)
	return nil
}

// GetProductSystemId ProductSystemId Getter
func (r AlibabaDamaiTickletQrcodeDecodeAPIRequest) GetProductSystemId() string {
	return r._productSystemId
}

var poolAlibabaDamaiTickletQrcodeDecodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiTickletQrcodeDecodeRequest()
	},
}

// GetAlibabaDamaiTickletQrcodeDecodeRequest 从 sync.Pool 获取 AlibabaDamaiTickletQrcodeDecodeAPIRequest
func GetAlibabaDamaiTickletQrcodeDecodeAPIRequest() *AlibabaDamaiTickletQrcodeDecodeAPIRequest {
	return poolAlibabaDamaiTickletQrcodeDecodeAPIRequest.Get().(*AlibabaDamaiTickletQrcodeDecodeAPIRequest)
}

// ReleaseAlibabaDamaiTickletQrcodeDecodeAPIRequest 将 AlibabaDamaiTickletQrcodeDecodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiTickletQrcodeDecodeAPIRequest(v *AlibabaDamaiTickletQrcodeDecodeAPIRequest) {
	v.Reset()
	poolAlibabaDamaiTickletQrcodeDecodeAPIRequest.Put(v)
}
