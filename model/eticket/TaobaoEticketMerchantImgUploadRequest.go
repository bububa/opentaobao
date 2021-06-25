package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商上传二维码图片 APIRequest
taobao.eticket.merchant.img.upload

电子凭证的码商可以通过这个接口，上传二维码图片
*/
type TaobaoEticketMerchantImgUploadRequest struct {
    model.Params

    // 二维码图片
    imgBytes   []byte 

}

func NewTaobaoEticketMerchantImgUploadRequest() *TaobaoEticketMerchantImgUploadRequest{
    return &TaobaoEticketMerchantImgUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantImgUploadRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.img.upload"
}

func (r TaobaoEticketMerchantImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantImgUploadRequest) SetImgBytes(imgBytes []byte) error {
    r.imgBytes = imgBytes
    r.Set("img_bytes", imgBytes)
    return nil
}

func (r TaobaoEticketMerchantImgUploadRequest) GetImgBytes() []byte {
    return r.imgBytes
}

