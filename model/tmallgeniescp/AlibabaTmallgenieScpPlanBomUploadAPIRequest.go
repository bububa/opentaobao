package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanBomUploadAPIRequest 计划BOM同步 API请求
// alibaba.tmallgenie.scp.plan.bom.upload
//
// 计划BOM同步
type AlibabaTmallgenieScpPlanBomUploadAPIRequest struct {
	model.Params
	// 对象
	_pbomRequest *AbstractRequest
}

// NewAlibabaTmallgenieScpPlanBomUploadRequest 初始化AlibabaTmallgenieScpPlanBomUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanBomUploadRequest() *AlibabaTmallgenieScpPlanBomUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanBomUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanBomUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.bom.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanBomUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PbomRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanBomUploadAPIRequest) SetPbomRequest(_pbomRequest *AbstractRequest) error {
	r._pbomRequest = _pbomRequest
	r.Set("pbom_request", _pbomRequest)
	return nil
}

// Get PbomRequest Getter
func (r AlibabaTmallgenieScpPlanBomUploadAPIRequest) GetPbomRequest() *AbstractRequest {
	return r._pbomRequest
}
