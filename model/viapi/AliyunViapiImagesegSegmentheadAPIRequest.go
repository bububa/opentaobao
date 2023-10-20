package viapi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmentheadAPIRequest 头像分割 API请求
// aliyun.viapi.imageseg.segmenthead
//
// 输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmentheadAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiImagesegSegmentheadRequest 初始化AliyunViapiImagesegSegmentheadAPIRequest对象
func NewAliyunViapiImagesegSegmentheadRequest() *AliyunViapiImagesegSegmentheadAPIRequest {
	return &AliyunViapiImagesegSegmentheadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunViapiImagesegSegmentheadAPIRequest) Reset() {
	r._imageUrl = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmentheadAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segmenthead"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmentheadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiImagesegSegmentheadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmentheadAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiImagesegSegmentheadAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

var poolAliyunViapiImagesegSegmentheadAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunViapiImagesegSegmentheadRequest()
	},
}

// GetAliyunViapiImagesegSegmentheadRequest 从 sync.Pool 获取 AliyunViapiImagesegSegmentheadAPIRequest
func GetAliyunViapiImagesegSegmentheadAPIRequest() *AliyunViapiImagesegSegmentheadAPIRequest {
	return poolAliyunViapiImagesegSegmentheadAPIRequest.Get().(*AliyunViapiImagesegSegmentheadAPIRequest)
}

// ReleaseAliyunViapiImagesegSegmentheadAPIRequest 将 AliyunViapiImagesegSegmentheadAPIRequest 放入 sync.Pool
func ReleaseAliyunViapiImagesegSegmentheadAPIRequest(v *AliyunViapiImagesegSegmentheadAPIRequest) {
	v.Reset()
	poolAliyunViapiImagesegSegmentheadAPIRequest.Put(v)
}
