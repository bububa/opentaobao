package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapifacebodycomparefaceAPIRequest 人脸比对1：1 API请求
// aliyun.viapi.facebody.compareface
//
// 基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapifacebodycomparefaceAPIRequest struct {
	model.Params
	// 图片url地址(http/https)
	_imageUrlA string
	// 图片url地址(http/https)
	_imageUrlB string
	// 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
	_imageType int64
}

// NewAliyunviapifacebodycomparefaceRequest 初始化AliyunviapifacebodycomparefaceAPIRequest对象
func NewAliyunviapifacebodycomparefaceRequest() *AliyunviapifacebodycomparefaceAPIRequest {
	return &AliyunviapifacebodycomparefaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapifacebodycomparefaceAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.facebody.compareface"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapifacebodycomparefaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapifacebodycomparefaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrlA is ImageUrlA Setter
// 图片url地址(http/https)
func (r *AliyunviapifacebodycomparefaceAPIRequest) SetImageUrlA(_imageUrlA string) error {
	r._imageUrlA = _imageUrlA
	r.Set("image_url_a", _imageUrlA)
	return nil
}

// GetImageUrlA ImageUrlA Getter
func (r AliyunviapifacebodycomparefaceAPIRequest) GetImageUrlA() string {
	return r._imageUrlA
}

// SetImageUrlB is ImageUrlB Setter
// 图片url地址(http/https)
func (r *AliyunviapifacebodycomparefaceAPIRequest) SetImageUrlB(_imageUrlB string) error {
	r._imageUrlB = _imageUrlB
	r.Set("image_url_b", _imageUrlB)
	return nil
}

// GetImageUrlB ImageUrlB Getter
func (r AliyunviapifacebodycomparefaceAPIRequest) GetImageUrlB() string {
	return r._imageUrlB
}

// SetImageType is ImageType Setter
// 图片类型, ,取值范围[0:图片URL, 1:图片Base64数据]
func (r *AliyunviapifacebodycomparefaceAPIRequest) SetImageType(_imageType int64) error {
	r._imageType = _imageType
	r.Set("image_type", _imageType)
	return nil
}

// GetImageType ImageType Getter
func (r AliyunviapifacebodycomparefaceAPIRequest) GetImageType() int64 {
	return r._imageType
}
