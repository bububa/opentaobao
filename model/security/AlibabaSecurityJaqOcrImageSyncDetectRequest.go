package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全图文识别同步检测接口 APIRequest
alibaba.security.jaq.ocr.image.sync.detect

图像字符识别同步检测接口
*/
type AlibabaSecurityJaqOcrImageSyncDetectRequest struct {
    model.Params

    // 待检测图像链接
    imageUrl   string 

}

func NewAlibabaSecurityJaqOcrImageSyncDetectRequest() *AlibabaSecurityJaqOcrImageSyncDetectRequest{
    return &AlibabaSecurityJaqOcrImageSyncDetectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqOcrImageSyncDetectRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.ocr.image.sync.detect"
}

func (r AlibabaSecurityJaqOcrImageSyncDetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqOcrImageSyncDetectRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

func (r AlibabaSecurityJaqOcrImageSyncDetectRequest) GetImageUrl() string {
    return r.imageUrl
}

