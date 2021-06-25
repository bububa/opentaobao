package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全智能鉴黄同步检测接口 APIRequest
alibaba.security.jaq.porn.image.sync.detect

同步黄图图像检测接口
*/
type AlibabaSecurityJaqPornImageSyncDetectRequest struct {
    model.Params

    // 待检测图片链接
    imageUrl   string 

}

func NewAlibabaSecurityJaqPornImageSyncDetectRequest() *AlibabaSecurityJaqPornImageSyncDetectRequest{
    return &AlibabaSecurityJaqPornImageSyncDetectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqPornImageSyncDetectRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.porn.image.sync.detect"
}

func (r AlibabaSecurityJaqPornImageSyncDetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqPornImageSyncDetectRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

func (r AlibabaSecurityJaqPornImageSyncDetectRequest) GetImageUrl() string {
    return r.imageUrl
}

