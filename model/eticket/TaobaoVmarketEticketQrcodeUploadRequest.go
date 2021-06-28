package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商二维码图片上传 APIRequest
taobao.vmarket.eticket.qrcode.upload

电子凭证的码商可以通过这个接口，上传他们发送的二维码图片
*/
type TaobaoVmarketEticketQrcodeUploadRequest struct {
    model.Params

    // 码商ID
    codeMerchantId   int64 

    // 上传的图片byte[]  小于300K，图片尺寸400*400以内
    imgBytes   []*model.File 

}

func NewTaobaoVmarketEticketQrcodeUploadRequest() *TaobaoVmarketEticketQrcodeUploadRequest{
    return &TaobaoVmarketEticketQrcodeUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketQrcodeUploadRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.qrcode.upload"
}

func (r TaobaoVmarketEticketQrcodeUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketQrcodeUploadRequest) SetCodeMerchantId(codeMerchantId int64) error {
    r.codeMerchantId = codeMerchantId
    r.Set("code_merchant_id", codeMerchantId)
    return nil
}

func (r TaobaoVmarketEticketQrcodeUploadRequest) GetCodeMerchantId() int64 {
    return r.codeMerchantId
}

func (r *TaobaoVmarketEticketQrcodeUploadRequest) SetImgBytes(imgBytes []*model.File) error {
    r.imgBytes = imgBytes
    r.Set("img_bytes", imgBytes)
    return nil
}

func (r TaobaoVmarketEticketQrcodeUploadRequest) GetImgBytes() []*model.File {
    return r.imgBytes
}

