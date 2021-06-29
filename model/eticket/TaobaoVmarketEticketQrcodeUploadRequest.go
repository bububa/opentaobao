package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商二维码图片上传 API请求
taobao.vmarket.eticket.qrcode.upload

电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
*/
type TaobaoVmarketEticketQrcodeUploadRequest struct {
    model.Params
    // 码商ID
    _codeMerchantId   int64
    // 上传的图片byte[]  小于300K，图片尺寸400*400以内
    _imgBytes   *model.File
}

// 初始化TaobaoVmarketEticketQrcodeUploadRequest对象
func NewTaobaoVmarketEticketQrcodeUploadRequest() *TaobaoVmarketEticketQrcodeUploadRequest{
    return &TaobaoVmarketEticketQrcodeUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketQrcodeUploadRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.qrcode.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketQrcodeUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CodeMerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketQrcodeUploadRequest) SetCodeMerchantId(_codeMerchantId int64) error {
    r._codeMerchantId = _codeMerchantId
    r.Set("code_merchant_id", _codeMerchantId)
    return nil
}

// CodeMerchantId Getter
func (r TaobaoVmarketEticketQrcodeUploadRequest) GetCodeMerchantId() int64 {
    return r._codeMerchantId
}
// ImgBytes Setter
// 上传的图片byte[]  小于300K，图片尺寸400*400以内
func (r *TaobaoVmarketEticketQrcodeUploadRequest) SetImgBytes(_imgBytes *model.File) error {
    r._imgBytes = _imgBytes
    r.Set("img_bytes", _imgBytes)
    return nil
}

// ImgBytes Getter
func (r TaobaoVmarketEticketQrcodeUploadRequest) GetImgBytes() *model.File {
    return r._imgBytes
}
