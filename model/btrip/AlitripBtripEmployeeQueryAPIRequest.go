package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripemployeequeryAPIRequest 企业员工查询 API请求
// alitrip.btrip.employee.query
//
// 企业员工查询
type AlitripbtripemployeequeryAPIRequest struct {
	model.Params
	// 入参对象。
	_paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest
}

// NewAlitripbtripemployeequeryRequest 初始化AlitripbtripemployeequeryAPIRequest对象
func NewAlitripbtripemployeequeryRequest() *AlitripbtripemployeequeryAPIRequest {
	return &AlitripbtripemployeequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripemployeequeryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.employee.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripemployeequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripemployeequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOpenEmployeeQueryRequest is ParamOpenEmployeeQueryRequest Setter
// 入参对象。
func (r *AlitripbtripemployeequeryAPIRequest) SetParamOpenEmployeeQueryRequest(_paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest) error {
	r._paramOpenEmployeeQueryRequest = _paramOpenEmployeeQueryRequest
	r.Set("param_open_employee_query_request", _paramOpenEmployeeQueryRequest)
	return nil
}

// GetParamOpenEmployeeQueryRequest ParamOpenEmployeeQueryRequest Getter
func (r AlitripbtripemployeequeryAPIRequest) GetParamOpenEmployeeQueryRequest() *OpenEmployeeQueryRequest {
	return r._paramOpenEmployeeQueryRequest
}
