package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalCommithandovercontentUpdateAPIRequest 修改已经提交的交接单 API请求
// cainiao.global.commithandovercontent.update
//
// 修改已经提交的交接单
type CainiaoGlobalCommithandovercontentUpdateAPIRequest struct {
	model.Params
	// 修改参数对象
	_openHandoverContentUpdateCommitRequest *OpenHandoverContentUpdateCommitRequest
}

// NewCainiaoGlobalCommithandovercontentUpdateRequest 初始化CainiaoGlobalCommithandovercontentUpdateAPIRequest对象
func NewCainiaoGlobalCommithandovercontentUpdateRequest() *CainiaoGlobalCommithandovercontentUpdateAPIRequest {
	return &CainiaoGlobalCommithandovercontentUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalCommithandovercontentUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.commithandovercontent.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalCommithandovercontentUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalCommithandovercontentUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenHandoverContentUpdateCommitRequest is OpenHandoverContentUpdateCommitRequest Setter
// 修改参数对象
func (r *CainiaoGlobalCommithandovercontentUpdateAPIRequest) SetOpenHandoverContentUpdateCommitRequest(_openHandoverContentUpdateCommitRequest *OpenHandoverContentUpdateCommitRequest) error {
	r._openHandoverContentUpdateCommitRequest = _openHandoverContentUpdateCommitRequest
	r.Set("open_handover_content_update_commit_request", _openHandoverContentUpdateCommitRequest)
	return nil
}

// GetOpenHandoverContentUpdateCommitRequest OpenHandoverContentUpdateCommitRequest Getter
func (r CainiaoGlobalCommithandovercontentUpdateAPIRequest) GetOpenHandoverContentUpdateCommitRequest() *OpenHandoverContentUpdateCommitRequest {
	return r._openHandoverContentUpdateCommitRequest
}
