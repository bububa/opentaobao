package viapi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmenthdbodyAPIRequest 高清人体分割 API请求
// aliyun.viapi.imageseg.segmenthdbody
//
// 对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmenthdbodyAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiImagesegSegmenthdbodyRequest 初始化AliyunViapiImagesegSegmenthdbodyAPIRequest对象
func NewAliyunViapiImagesegSegmenthdbodyRequest() *AliyunViapiImagesegSegmenthdbodyAPIRequest {
	return &AliyunViapiImagesegSegmenthdbodyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunViapiImagesegSegmenthdbodyAPIRequest) Reset() {
	r._imageUrl = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segmenthdbody"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiImagesegSegmenthdbodyAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiImagesegSegmenthdbodyAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

var poolAliyunViapiImagesegSegmenthdbodyAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunViapiImagesegSegmenthdbodyRequest()
	},
}

// GetAliyunViapiImagesegSegmenthdbodyRequest 从 sync.Pool 获取 AliyunViapiImagesegSegmenthdbodyAPIRequest
func GetAliyunViapiImagesegSegmenthdbodyAPIRequest() *AliyunViapiImagesegSegmenthdbodyAPIRequest {
	return poolAliyunViapiImagesegSegmenthdbodyAPIRequest.Get().(*AliyunViapiImagesegSegmenthdbodyAPIRequest)
}

// ReleaseAliyunViapiImagesegSegmenthdbodyAPIRequest 将 AliyunViapiImagesegSegmenthdbodyAPIRequest 放入 sync.Pool
func ReleaseAliyunViapiImagesegSegmenthdbodyAPIRequest(v *AliyunViapiImagesegSegmenthdbodyAPIRequest) {
	v.Reset()
	poolAliyunViapiImagesegSegmenthdbodyAPIRequest.Put(v)
}
