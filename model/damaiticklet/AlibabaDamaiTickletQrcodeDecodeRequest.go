package damaiticklet

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
票夹-动态二维码-解码 API请求
alibaba.damai.ticklet.qrcode.decode

对于票夹的动态二维码进行解码
*/
type AlibabaDamaiTickletQrcodeDecodeRequest struct {
    model.Params
    // 生产系统
    _productSystemId   string
    // 加密二维码
    _encryptedQrCode   string
}

// 初始化AlibabaDamaiTickletQrcodeDecodeRequest对象
func NewAlibabaDamaiTickletQrcodeDecodeRequest() *AlibabaDamaiTickletQrcodeDecodeRequest{
    return &AlibabaDamaiTickletQrcodeDecodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetApiMethodName() string {
    return "alibaba.damai.ticklet.qrcode.decode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductSystemId Setter
// 生产系统
func (r *AlibabaDamaiTickletQrcodeDecodeRequest) SetProductSystemId(_productSystemId string) error {
    r._productSystemId = _productSystemId
    r.Set("product_system_id", _productSystemId)
    return nil
}

// ProductSystemId Getter
func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetProductSystemId() string {
    return r._productSystemId
}
// EncryptedQrCode Setter
// 加密二维码
func (r *AlibabaDamaiTickletQrcodeDecodeRequest) SetEncryptedQrCode(_encryptedQrCode string) error {
    r._encryptedQrCode = _encryptedQrCode
    r.Set("encrypted_qr_code", _encryptedQrCode)
    return nil
}

// EncryptedQrCode Getter
func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetEncryptedQrCode() string {
    return r._encryptedQrCode
}
