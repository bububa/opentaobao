package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品分割 API请求
aliyun.viapi.imageseg.segmentcomodity

识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiImagesegSegmentcomodityRequest struct {
    model.Params
    // 待检测图片链接
    _imageUrl   string
}

// 初始化AliyunViapiImagesegSegmentcomodityRequest对象
func NewAliyunViapiImagesegSegmentcomodityRequest() *AliyunViapiImagesegSegmentcomodityRequest{
    return &AliyunViapiImagesegSegmentcomodityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmentcomodityRequest) GetApiMethodName() string {
    return "aliyun.viapi.imageseg.segmentcomodity"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmentcomodityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmentcomodityRequest) SetImageUrl(_imageUrl string) error {
    r._imageUrl = _imageUrl
    r.Set("image_url", _imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiImagesegSegmentcomodityRequest) GetImageUrl() string {
    return r._imageUrl
}
