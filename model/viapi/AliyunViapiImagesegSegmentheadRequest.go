package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
头像分割 API请求
aliyun.viapi.imageseg.segmenthead

输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImagesegSegmentheadRequest struct {
    model.Params
    // 待检测图片链接
    imageUrl   string
}

// 初始化AliyunViapiImagesegSegmentheadRequest对象
func NewAliyunViapiImagesegSegmentheadRequest() *AliyunViapiImagesegSegmentheadRequest{
    return &AliyunViapiImagesegSegmentheadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmentheadRequest) GetApiMethodName() string {
    return "aliyun.viapi.imageseg.segmenthead"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmentheadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmentheadRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiImagesegSegmentheadRequest) GetImageUrl() string {
    return r.imageUrl
}
