package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定流程上下文参数 API请求
alibaba.mosflow.work.queryvariables

业务查询指定流程上下文内容
*/
type AlibabaMosflowWorkQueryvariablesRequest struct {
    model.Params
    // 流程实例ID
    _processInstanceId   string
}

// 初始化AlibabaMosflowWorkQueryvariablesRequest对象
func NewAlibabaMosflowWorkQueryvariablesRequest() *AlibabaMosflowWorkQueryvariablesRequest{
    return &AlibabaMosflowWorkQueryvariablesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosflowWorkQueryvariablesRequest) GetApiMethodName() string {
    return "alibaba.mosflow.work.queryvariables"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosflowWorkQueryvariablesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProcessInstanceId Setter
// 流程实例ID
func (r *AlibabaMosflowWorkQueryvariablesRequest) SetProcessInstanceId(_processInstanceId string) error {
    r._processInstanceId = _processInstanceId
    r.Set("process_instance_id", _processInstanceId)
    return nil
}

// ProcessInstanceId Getter
func (r AlibabaMosflowWorkQueryvariablesRequest) GetProcessInstanceId() string {
    return r._processInstanceId
}
