package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfoincominginvoicepytimageuploadAPIRequest 票易通发票影像上传 API请求
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
type AlibabacfoincominginvoicepytimageuploadAPIRequest struct {
	model.Params
	// 影像上传实体
	_uploadRequest *ImageUploadRequest
}

// NewAlibabacfoincominginvoicepytimageuploadRequest 初始化AlibabacfoincominginvoicepytimageuploadAPIRequest对象
func NewAlibabacfoincominginvoicepytimageuploadRequest() *AlibabacfoincominginvoicepytimageuploadAPIRequest {
	return &AlibabacfoincominginvoicepytimageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacfoincominginvoicepytimageuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.cfo.incoming.invoice.pyt.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacfoincominginvoicepytimageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacfoincominginvoicepytimageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadRequest is UploadRequest Setter
// 影像上传实体
func (r *AlibabacfoincominginvoicepytimageuploadAPIRequest) SetUploadRequest(_uploadRequest *ImageUploadRequest) error {
	r._uploadRequest = _uploadRequest
	r.Set("upload_request", _uploadRequest)
	return nil
}

// GetUploadRequest UploadRequest Getter
func (r AlibabacfoincominginvoicepytimageuploadAPIRequest) GetUploadRequest() *ImageUploadRequest {
	return r._uploadRequest
}
