package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商上传二维码图片 API请求
taobao.eticket.merchant.img.upload

电子凭证的码商可以通过这个接口，上传二维码图片
*/
type TaobaoEticketMerchantImgUploadRequest struct {
    model.Params
    // 二维码图片
    imgBytes   []*model.File
}

// 初始化TaobaoEticketMerchantImgUploadRequest对象
func NewTaobaoEticketMerchantImgUploadRequest() *TaobaoEticketMerchantImgUploadRequest{
    return &TaobaoEticketMerchantImgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantImgUploadRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.img.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImgBytes Setter
// 二维码图片
func (r *TaobaoEticketMerchantImgUploadRequest) SetImgBytes(imgBytes []*model.File) error {
    r.imgBytes = imgBytes
    r.Set("img_bytes", imgBytes)
    return nil
}

// ImgBytes Getter
func (r TaobaoEticketMerchantImgUploadRequest) GetImgBytes() []*model.File {
    return r.imgBytes
}
