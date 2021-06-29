package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全图文识别同步检测接口 API请求
alibaba.security.jaq.ocr.image.sync.detect

图像字符识别同步检测接口
*/
type AlibabaSecurityJaqOcrImageSyncDetectRequest struct {
    model.Params
    // 待检测图像链接
    _imageUrl   string
}

// 初始化AlibabaSecurityJaqOcrImageSyncDetectRequest对象
func NewAlibabaSecurityJaqOcrImageSyncDetectRequest() *AlibabaSecurityJaqOcrImageSyncDetectRequest{
    return &AlibabaSecurityJaqOcrImageSyncDetectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqOcrImageSyncDetectRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.ocr.image.sync.detect"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqOcrImageSyncDetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图像链接
func (r *AlibabaSecurityJaqOcrImageSyncDetectRequest) SetImageUrl(_imageUrl string) error {
    r._imageUrl = _imageUrl
    r.Set("image_url", _imageUrl)
    return nil
}

// ImageUrl Getter
func (r AlibabaSecurityJaqOcrImageSyncDetectRequest) GetImageUrl() string {
    return r._imageUrl
}
