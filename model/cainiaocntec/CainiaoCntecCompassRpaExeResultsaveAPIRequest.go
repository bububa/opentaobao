package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecCompassRpaExeResultsaveAPIRequest rpa执行结果回传 API请求
// cainiao.cntec.compass.rpa.exe.resultsave
//
// rpa执行结果回传
type CainiaoCntecCompassRpaExeResultsaveAPIRequest struct {
	model.Params
	// 请求参数
	_rpaExeResultByUuidReq *RpaExeResultByUuidReq
}

// NewCainiaoCntecCompassRpaExeResultsaveRequest 初始化CainiaoCntecCompassRpaExeResultsaveAPIRequest对象
func NewCainiaoCntecCompassRpaExeResultsaveRequest() *CainiaoCntecCompassRpaExeResultsaveAPIRequest {
	return &CainiaoCntecCompassRpaExeResultsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecCompassRpaExeResultsaveAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.compass.rpa.exe.resultsave"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecCompassRpaExeResultsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntecCompassRpaExeResultsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRpaExeResultByUuidReq is RpaExeResultByUuidReq Setter
// 请求参数
func (r *CainiaoCntecCompassRpaExeResultsaveAPIRequest) SetRpaExeResultByUuidReq(_rpaExeResultByUuidReq *RpaExeResultByUuidReq) error {
	r._rpaExeResultByUuidReq = _rpaExeResultByUuidReq
	r.Set("rpa_exe_result_by_uuid_req", _rpaExeResultByUuidReq)
	return nil
}

// GetRpaExeResultByUuidReq RpaExeResultByUuidReq Getter
func (r CainiaoCntecCompassRpaExeResultsaveAPIRequest) GetRpaExeResultByUuidReq() *RpaExeResultByUuidReq {
	return r._rpaExeResultByUuidReq
}
