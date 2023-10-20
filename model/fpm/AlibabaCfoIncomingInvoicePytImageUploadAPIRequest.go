package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfoIncomingInvoicePytImageUploadAPIRequest 票易通发票影像上传 API请求
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
type AlibabaCfoIncomingInvoicePytImageUploadAPIRequest struct {
	model.Params
	// 影像上传实体
	_uploadRequest *ImageUploadRequest
}

// NewAlibabaCfoIncomingInvoicePytImageUploadRequest 初始化AlibabaCfoIncomingInvoicePytImageUploadAPIRequest对象
func NewAlibabaCfoIncomingInvoicePytImageUploadRequest() *AlibabaCfoIncomingInvoicePytImageUploadAPIRequest {
	return &AlibabaCfoIncomingInvoicePytImageUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.cfo.incoming.invoice.pyt.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUploadRequest is UploadRequest Setter
// 影像上传实体
func (r *AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) SetUploadRequest(_uploadRequest *ImageUploadRequest) error {
	r._uploadRequest = _uploadRequest
	r.Set("upload_request", _uploadRequest)
	return nil
}

// GetUploadRequest UploadRequest Getter
func (r AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) GetUploadRequest() *ImageUploadRequest {
	return r._uploadRequest
}
