package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqOcrImageSyncDetectAPIRequest 聚安全图文识别同步检测接口 API请求
// alibaba.security.jaq.ocr.image.sync.detect
//
// 图像字符识别同步检测接口
type AlibabaSecurityJaqOcrImageSyncDetectAPIRequest struct {
	model.Params
	// 待检测图像链接
	_imageUrl string
}

// NewAlibabaSecurityJaqOcrImageSyncDetectRequest 初始化AlibabaSecurityJaqOcrImageSyncDetectAPIRequest对象
func NewAlibabaSecurityJaqOcrImageSyncDetectRequest() *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest {
	return &AlibabaSecurityJaqOcrImageSyncDetectAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) Reset() {
	r._imageUrl = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.ocr.image.sync.detect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图像链接
func (r *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

var poolAlibabaSecurityJaqOcrImageSyncDetectAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqOcrImageSyncDetectRequest()
	},
}

// GetAlibabaSecurityJaqOcrImageSyncDetectRequest 从 sync.Pool 获取 AlibabaSecurityJaqOcrImageSyncDetectAPIRequest
func GetAlibabaSecurityJaqOcrImageSyncDetectAPIRequest() *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest {
	return poolAlibabaSecurityJaqOcrImageSyncDetectAPIRequest.Get().(*AlibabaSecurityJaqOcrImageSyncDetectAPIRequest)
}

// ReleaseAlibabaSecurityJaqOcrImageSyncDetectAPIRequest 将 AlibabaSecurityJaqOcrImageSyncDetectAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqOcrImageSyncDetectAPIRequest(v *AlibabaSecurityJaqOcrImageSyncDetectAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqOcrImageSyncDetectAPIRequest.Put(v)
}
