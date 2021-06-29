package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全智能鉴黄同步检测接口 API请求
alibaba.security.jaq.porn.image.sync.detect

同步黄图图像检测接口
*/
type AlibabaSecurityJaqPornImageSyncDetectRequest struct {
    model.Params
    // 待检测图片链接
    _imageUrl   string
}

// 初始化AlibabaSecurityJaqPornImageSyncDetectRequest对象
func NewAlibabaSecurityJaqPornImageSyncDetectRequest() *AlibabaSecurityJaqPornImageSyncDetectRequest{
    return &AlibabaSecurityJaqPornImageSyncDetectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqPornImageSyncDetectRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.porn.image.sync.detect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqPornImageSyncDetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AlibabaSecurityJaqPornImageSyncDetectRequest) SetImageUrl(_imageUrl string) error {
    r._imageUrl = _imageUrl
    r.Set("image_url", _imageUrl)
    return nil
}

// ImageUrl Getter
func (r AlibabaSecurityJaqPornImageSyncDetectRequest) GetImageUrl() string {
    return r._imageUrl
}
