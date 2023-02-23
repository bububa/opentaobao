package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest 静态二维码绑定 API请求
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
type AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest struct {
	model.Params
	// 设备唯一标识，如mac或sn
	_identifier string
	// 产品唯一标识，星空平台产品pk
	_productKey string
	// 签名摘要
	_digest string
}

// NewAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindRequest 初始化AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindRequest() *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifier is Identifier Setter
// 设备唯一标识，如mac或sn
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetProductKey is ProductKey Setter
// 产品唯一标识，星空平台产品pk
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) SetProductKey(_productKey string) error {
	r._productKey = _productKey
	r.Set("product_key", _productKey)
	return nil
}

// GetProductKey ProductKey Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) GetProductKey() string {
	return r._productKey
}

// SetDigest is Digest Setter
// 签名摘要
func (r *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) SetDigest(_digest string) error {
	r._digest = _digest
	r.Set("digest", _digest)
	return nil
}

// GetDigest Digest Getter
func (r AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIRequest) GetDigest() string {
	return r._digest
}
