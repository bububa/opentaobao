package fpm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafpmfileuploadAPIRequest 结算单文件上传 API请求
// alibaba.fpm.file.upload
//
// 结算单文件上传
type AlibabafpmfileuploadAPIRequest struct {
	model.Params
	// 实体
	_bizDto *FileUploadRequestDto
}

// NewAlibabafpmfileuploadRequest 初始化AlibabafpmfileuploadAPIRequest对象
func NewAlibabafpmfileuploadRequest() *AlibabafpmfileuploadAPIRequest {
	return &AlibabafpmfileuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafpmfileuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.fpm.file.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafpmfileuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafpmfileuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizDto is BizDto Setter
// 实体
func (r *AlibabafpmfileuploadAPIRequest) SetBizDto(_bizDto *FileUploadRequestDto) error {
	r._bizDto = _bizDto
	r.Set("biz_dto", _bizDto)
	return nil
}

// GetBizDto BizDto Getter
func (r AlibabafpmfileuploadAPIRequest) GetBizDto() *FileUploadRequestDto {
	return r._bizDto
}
