package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
高清人体分割 API请求
aliyun.viapi.imageseg.segmenthdbody

对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImagesegSegmenthdbodyAPIRequest struct {
    model.Params
    // 待检测图片链接
    _imageUrl   string
}

// 初始化AliyunViapiImagesegSegmenthdbodyAPIRequest对象
func NewAliyunViapiImagesegSegmenthdbodyRequest() *AliyunViapiImagesegSegmenthdbodyAPIRequest{
    return &AliyunViapiImagesegSegmenthdbodyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetApiMethodName() string {
    return "aliyun.viapi.imageseg.segmenthdbody"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmenthdbodyAPIRequest) SetImageUrl(_imageUrl string) error {
    r._imageUrl = _imageUrl
    r.Set("image_url", _imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetImageUrl() string {
    return r._imageUrl
}
