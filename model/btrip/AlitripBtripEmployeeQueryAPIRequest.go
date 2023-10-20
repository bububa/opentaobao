package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripEmployeeQueryAPIRequest 企业员工查询 API请求
// alitrip.btrip.employee.query
//
// 企业员工查询
type AlitripBtripEmployeeQueryAPIRequest struct {
	model.Params
	// 入参对象。
	_paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest
}

// NewAlitripBtripEmployeeQueryRequest 初始化AlitripBtripEmployeeQueryAPIRequest对象
func NewAlitripBtripEmployeeQueryRequest() *AlitripBtripEmployeeQueryAPIRequest {
	return &AlitripBtripEmployeeQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripEmployeeQueryAPIRequest) Reset() {
	r._paramOpenEmployeeQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripEmployeeQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.employee.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripEmployeeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripEmployeeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOpenEmployeeQueryRequest is ParamOpenEmployeeQueryRequest Setter
// 入参对象。
func (r *AlitripBtripEmployeeQueryAPIRequest) SetParamOpenEmployeeQueryRequest(_paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest) error {
	r._paramOpenEmployeeQueryRequest = _paramOpenEmployeeQueryRequest
	r.Set("param_open_employee_query_request", _paramOpenEmployeeQueryRequest)
	return nil
}

// GetParamOpenEmployeeQueryRequest ParamOpenEmployeeQueryRequest Getter
func (r AlitripBtripEmployeeQueryAPIRequest) GetParamOpenEmployeeQueryRequest() *OpenEmployeeQueryRequest {
	return r._paramOpenEmployeeQueryRequest
}

var poolAlitripBtripEmployeeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripEmployeeQueryRequest()
	},
}

// GetAlitripBtripEmployeeQueryRequest 从 sync.Pool 获取 AlitripBtripEmployeeQueryAPIRequest
func GetAlitripBtripEmployeeQueryAPIRequest() *AlitripBtripEmployeeQueryAPIRequest {
	return poolAlitripBtripEmployeeQueryAPIRequest.Get().(*AlitripBtripEmployeeQueryAPIRequest)
}

// ReleaseAlitripBtripEmployeeQueryAPIRequest 将 AlitripBtripEmployeeQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripEmployeeQueryAPIRequest(v *AlitripBtripEmployeeQueryAPIRequest) {
	v.Reset()
	poolAlitripBtripEmployeeQueryAPIRequest.Put(v)
}
