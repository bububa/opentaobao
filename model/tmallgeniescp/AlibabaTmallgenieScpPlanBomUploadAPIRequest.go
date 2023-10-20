package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanbomuploadAPIRequest 计划BOM同步 API请求
// alibaba.tmallgenie.scp.plan.bom.upload
//
// 计划BOM同步
type AlibabatmallgeniescpplanbomuploadAPIRequest struct {
	model.Params
	// 对象
	_pbomRequest *AbstractRequest
}

// NewAlibabatmallgeniescpplanbomuploadRequest 初始化AlibabatmallgeniescpplanbomuploadAPIRequest对象
func NewAlibabatmallgeniescpplanbomuploadRequest() *AlibabatmallgeniescpplanbomuploadAPIRequest {
	return &AlibabatmallgeniescpplanbomuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanbomuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.bom.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanbomuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanbomuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPbomRequest is PbomRequest Setter
// 对象
func (r *AlibabatmallgeniescpplanbomuploadAPIRequest) SetPbomRequest(_pbomRequest *AbstractRequest) error {
	r._pbomRequest = _pbomRequest
	r.Set("pbom_request", _pbomRequest)
	return nil
}

// GetPbomRequest PbomRequest Getter
func (r AlibabatmallgeniescpplanbomuploadAPIRequest) GetPbomRequest() *AbstractRequest {
	return r._pbomRequest
}
