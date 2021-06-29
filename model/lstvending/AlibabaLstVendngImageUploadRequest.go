package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
售货机商品图片上传 APIRequest
alibaba.lst.vendng.image.upload

零售通自动售货机商品图片上传接口，主要为ISV厂商提供图片同步的通道，从而建立统一的商品图片库。
*/
type AlibabaLstVendngImageUploadRequest struct {
    model.Params

    // 图片文件字节数组
    imgBytes   []*model.File 

}

func NewAlibabaLstVendngImageUploadRequest() *AlibabaLstVendngImageUploadRequest{
    return &AlibabaLstVendngImageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstVendngImageUploadRequest) GetApiMethodName() string {
    return "alibaba.lst.vendng.image.upload"
}

func (r AlibabaLstVendngImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstVendngImageUploadRequest) SetImgBytes(imgBytes []*model.File) error {
    r.imgBytes = imgBytes
    r.Set("img_bytes", imgBytes)
    return nil
}

func (r AlibabaLstVendngImageUploadRequest) GetImgBytes() []*model.File {
    return r.imgBytes
}

