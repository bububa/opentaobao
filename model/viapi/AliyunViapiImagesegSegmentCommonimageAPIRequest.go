package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmentCommonimageAPIRequest 通用分割 API请求
// aliyun.viapi.imageseg.segment.commonimage
//
// 识别输入图像中的视觉中心物体轮廓，与背景进行分离，返回分割后的前景物体图（4通道png透明图）。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmentCommonimageAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiImagesegSegmentCommonimageRequest 初始化AliyunViapiImagesegSegmentCommonimageAPIRequest对象
func NewAliyunViapiImagesegSegmentCommonimageRequest() *AliyunViapiImagesegSegmentCommonimageAPIRequest {
	return &AliyunViapiImagesegSegmentCommonimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmentCommonimageAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segment.commonimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmentCommonimageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmentCommonimageAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiImagesegSegmentCommonimageAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
