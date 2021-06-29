package damaiticklet

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
票夹-动态二维码-解码 APIRequest
alibaba.damai.ticklet.qrcode.decode

对于票夹的动态二维码进行解码
*/
type AlibabaDamaiTickletQrcodeDecodeRequest struct {
    model.Params

    // 生产系统
    productSystemId   string 

    // 加密二维码
    encryptedQrCode   string 

}

func NewAlibabaDamaiTickletQrcodeDecodeRequest() *AlibabaDamaiTickletQrcodeDecodeRequest{
    return &AlibabaDamaiTickletQrcodeDecodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetApiMethodName() string {
    return "alibaba.damai.ticklet.qrcode.decode"
}

func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiTickletQrcodeDecodeRequest) SetProductSystemId(productSystemId string) error {
    r.productSystemId = productSystemId
    r.Set("product_system_id", productSystemId)
    return nil
}

func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetProductSystemId() string {
    return r.productSystemId
}

func (r *AlibabaDamaiTickletQrcodeDecodeRequest) SetEncryptedQrCode(encryptedQrCode string) error {
    r.encryptedQrCode = encryptedQrCode
    r.Set("encrypted_qr_code", encryptedQrCode)
    return nil
}

func (r AlibabaDamaiTickletQrcodeDecodeRequest) GetEncryptedQrCode() string {
    return r.encryptedQrCode
}

