package fpm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFpmFileUploadAPIRequest 结算单文件上传 API请求
// alibaba.fpm.file.upload
//
// 结算单文件上传
type AlibabaFpmFileUploadAPIRequest struct {
	model.Params
	// 实体
	_bizDto *FileUploadRequestDto
}

// NewAlibabaFpmFileUploadRequest 初始化AlibabaFpmFileUploadAPIRequest对象
func NewAlibabaFpmFileUploadRequest() *AlibabaFpmFileUploadAPIRequest {
	return &AlibabaFpmFileUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFpmFileUploadAPIRequest) Reset() {
	r._bizDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFpmFileUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.fpm.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFpmFileUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFpmFileUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizDto is BizDto Setter
// 实体
func (r *AlibabaFpmFileUploadAPIRequest) SetBizDto(_bizDto *FileUploadRequestDto) error {
	r._bizDto = _bizDto
	r.Set("biz_dto", _bizDto)
	return nil
}

// GetBizDto BizDto Getter
func (r AlibabaFpmFileUploadAPIRequest) GetBizDto() *FileUploadRequestDto {
	return r._bizDto
}

var poolAlibabaFpmFileUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFpmFileUploadRequest()
	},
}

// GetAlibabaFpmFileUploadRequest 从 sync.Pool 获取 AlibabaFpmFileUploadAPIRequest
func GetAlibabaFpmFileUploadAPIRequest() *AlibabaFpmFileUploadAPIRequest {
	return poolAlibabaFpmFileUploadAPIRequest.Get().(*AlibabaFpmFileUploadAPIRequest)
}

// ReleaseAlibabaFpmFileUploadAPIRequest 将 AlibabaFpmFileUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaFpmFileUploadAPIRequest(v *AlibabaFpmFileUploadAPIRequest) {
	v.Reset()
	poolAlibabaFpmFileUploadAPIRequest.Put(v)
}
