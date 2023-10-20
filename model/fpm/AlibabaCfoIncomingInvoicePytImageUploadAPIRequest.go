package fpm

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) Reset() {
	r._uploadRequest = nil
	r.Params.ToZero()
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

var poolAlibabaCfoIncomingInvoicePytImageUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCfoIncomingInvoicePytImageUploadRequest()
	},
}

// GetAlibabaCfoIncomingInvoicePytImageUploadRequest 从 sync.Pool 获取 AlibabaCfoIncomingInvoicePytImageUploadAPIRequest
func GetAlibabaCfoIncomingInvoicePytImageUploadAPIRequest() *AlibabaCfoIncomingInvoicePytImageUploadAPIRequest {
	return poolAlibabaCfoIncomingInvoicePytImageUploadAPIRequest.Get().(*AlibabaCfoIncomingInvoicePytImageUploadAPIRequest)
}

// ReleaseAlibabaCfoIncomingInvoicePytImageUploadAPIRequest 将 AlibabaCfoIncomingInvoicePytImageUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaCfoIncomingInvoicePytImageUploadAPIRequest(v *AlibabaCfoIncomingInvoicePytImageUploadAPIRequest) {
	v.Reset()
	poolAlibabaCfoIncomingInvoicePytImageUploadAPIRequest.Put(v)
}
