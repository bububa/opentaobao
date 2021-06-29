package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人脸比对1：1 API请求
aliyun.viapi.facebody.compareface

基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiFacebodyComparefaceRequest struct {
    model.Params
    // 图片url地址(http/https)
    _imageUrlA   string
    // 图片url地址(http/https)
    _imageUrlB   string
    // 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
    _imageType   int64
}

// 初始化AliyunViapiFacebodyComparefaceRequest对象
func NewAliyunViapiFacebodyComparefaceRequest() *AliyunViapiFacebodyComparefaceRequest{
    return &AliyunViapiFacebodyComparefaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiFacebodyComparefaceRequest) GetApiMethodName() string {
    return "aliyun.viapi.facebody.compareface"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiFacebodyComparefaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrlA Setter
// 图片url地址(http/https)
func (r *AliyunViapiFacebodyComparefaceRequest) SetImageUrlA(_imageUrlA string) error {
    r._imageUrlA = _imageUrlA
    r.Set("image_url_a", _imageUrlA)
    return nil
}

// ImageUrlA Getter
func (r AliyunViapiFacebodyComparefaceRequest) GetImageUrlA() string {
    return r._imageUrlA
}
// ImageUrlB Setter
// 图片url地址(http/https)
func (r *AliyunViapiFacebodyComparefaceRequest) SetImageUrlB(_imageUrlB string) error {
    r._imageUrlB = _imageUrlB
    r.Set("image_url_b", _imageUrlB)
    return nil
}

// ImageUrlB Getter
func (r AliyunViapiFacebodyComparefaceRequest) GetImageUrlB() string {
    return r._imageUrlB
}
// ImageType Setter
// 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
func (r *AliyunViapiFacebodyComparefaceRequest) SetImageType(_imageType int64) error {
    r._imageType = _imageType
    r.Set("image_type", _imageType)
    return nil
}

// ImageType Getter
func (r AliyunViapiFacebodyComparefaceRequest) GetImageType() int64 {
    return r._imageType
}
