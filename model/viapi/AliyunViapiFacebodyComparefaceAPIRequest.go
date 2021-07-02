package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiFacebodyComparefaceAPIRequest 人脸比对1：1 API请求
// aliyun.viapi.facebody.compareface
//
// 基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiFacebodyComparefaceAPIRequest struct {
	model.Params
	// 图片url地址(http/https)
	_imageUrlA string
	// 图片url地址(http/https)
	_imageUrlB string
	// 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
	_imageType int64
}

// NewAliyunViapiFacebodyComparefaceRequest 初始化AliyunViapiFacebodyComparefaceAPIRequest对象
func NewAliyunViapiFacebodyComparefaceRequest() *AliyunViapiFacebodyComparefaceAPIRequest {
	return &AliyunViapiFacebodyComparefaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiFacebodyComparefaceAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.facebody.compareface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiFacebodyComparefaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImageUrlA Setter
// 图片url地址(http/https)
func (r *AliyunViapiFacebodyComparefaceAPIRequest) SetImageUrlA(_imageUrlA string) error {
	r._imageUrlA = _imageUrlA
	r.Set("image_url_a", _imageUrlA)
	return nil
}

// Get ImageUrlA Getter
func (r AliyunViapiFacebodyComparefaceAPIRequest) GetImageUrlA() string {
	return r._imageUrlA
}

// Set is ImageUrlB Setter
// 图片url地址(http/https)
func (r *AliyunViapiFacebodyComparefaceAPIRequest) SetImageUrlB(_imageUrlB string) error {
	r._imageUrlB = _imageUrlB
	r.Set("image_url_b", _imageUrlB)
	return nil
}

// Get ImageUrlB Getter
func (r AliyunViapiFacebodyComparefaceAPIRequest) GetImageUrlB() string {
	return r._imageUrlB
}

// Set is ImageType Setter
// 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
func (r *AliyunViapiFacebodyComparefaceAPIRequest) SetImageType(_imageType int64) error {
	r._imageType = _imageType
	r.Set("image_type", _imageType)
	return nil
}

// Get ImageType Getter
func (r AliyunViapiFacebodyComparefaceAPIRequest) GetImageType() int64 {
	return r._imageType
}
