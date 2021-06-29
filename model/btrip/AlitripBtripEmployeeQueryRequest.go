package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业员工查询 APIRequest
alitrip.btrip.employee.query

企业员工查询
*/
type AlitripBtripEmployeeQueryRequest struct {
    model.Params

    // 入参对象。
    paramOpenEmployeeQueryRequest   *OpenEmployeeQueryRequest 

}

func NewAlitripBtripEmployeeQueryRequest() *AlitripBtripEmployeeQueryRequest{
    return &AlitripBtripEmployeeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripEmployeeQueryRequest) GetApiMethodName() string {
    return "alitrip.btrip.employee.query"
}

func (r AlitripBtripEmployeeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripEmployeeQueryRequest) SetParamOpenEmployeeQueryRequest(paramOpenEmployeeQueryRequest *OpenEmployeeQueryRequest) error {
    r.paramOpenEmployeeQueryRequest = paramOpenEmployeeQueryRequest
    r.Set("param_open_employee_query_request", paramOpenEmployeeQueryRequest)
    return nil
}

func (r AlitripBtripEmployeeQueryRequest) GetParamOpenEmployeeQueryRequest() *OpenEmployeeQueryRequest {
    return r.paramOpenEmployeeQueryRequest
}

