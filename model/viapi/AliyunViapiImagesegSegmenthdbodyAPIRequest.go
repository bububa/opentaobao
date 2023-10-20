package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapiimagesegsegmenthdbodyAPIRequest 高清人体分割 API请求
// aliyun.viapi.imageseg.segmenthdbody
//
// 对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapiimagesegsegmenthdbodyAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunviapiimagesegsegmenthdbodyRequest 初始化AliyunviapiimagesegsegmenthdbodyAPIRequest对象
func NewAliyunviapiimagesegsegmenthdbodyRequest() *AliyunviapiimagesegsegmenthdbodyAPIRequest {
	return &AliyunviapiimagesegsegmenthdbodyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapiimagesegsegmenthdbodyAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.imageseg.segmenthdbody"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapiimagesegsegmenthdbodyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapiimagesegsegmenthdbodyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunviapiimagesegsegmenthdbodyAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunviapiimagesegsegmenthdbodyAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
