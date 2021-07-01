package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunViapiImagesegSegmentcomodityAPIRequest
商品分割 API请求
aliyun.viapi.imageseg.segmentcomodity

识别输入图像中的商品轮廓，与背景进行分离，返回分割后的前景商品图（4通道png透明图），适应单商品/多商品、复杂背景。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html ) */
type AliyunViapiImagesegSegmentcomodityAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiImagesegSegmentcomodityRequest 初始化AliyunViapiImagesegSegmentcomodityAPIRequest对象
func NewAliyunViapiImagesegSegmentcomodityRequest() *AliyunViapiImagesegSegmentcomodityAPIRequest {
	return &AliyunViapiImagesegSegmentcomodityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segmentcomodity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmentcomodityAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// Get ImageUrl Getter
func (r AliyunViapiImagesegSegmentcomodityAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
