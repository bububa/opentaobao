package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业员工查询 API请求
alitrip.btrip.employee.query

企业员工查询
*/
type AlitripBtripEmployeeQueryRequest struct {
    model.Params
    // 入参对象。
    _paramOpenEmployeeQueryRequest   *OpenEmployeeQueryRequest
}

// 初始化AlitripBtripEmployeeQueryRequest对象
func NewAlitripBtripEmployeeQueryRequest() *AlitripBtripEmployeeQueryRequest{
    return &AlitripBtripEmployeeQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripEmployeeQueryRequest) GetApiMethodName() string {
    return "alitrip.btrip.employee.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripEmployeeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOpenEmployeeQueryRequest Setter
// 入参对象。
func (r *AlitripBtripEmployeeQueryRequest) SetParamOpenEmployeeQueryRequest(_paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest) error {
    r._paramOpenEmployeeQueryRequest = _paramOpenEmployeeQueryRequest
    r.Set("param_open_employee_query_request", _paramOpenEmployeeQueryRequest)
    return nil
}

// ParamOpenEmployeeQueryRequest Getter
func (r AlitripBtripEmployeeQueryRequest) GetParamOpenEmployeeQueryRequest() *OpenEmployeeQueryRequest {
    return r._paramOpenEmployeeQueryRequest
}
