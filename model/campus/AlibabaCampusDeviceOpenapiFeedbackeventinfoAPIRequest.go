package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IVS事件处理反馈接口 API请求
alibaba.campus.device.openapi.feedbackeventinfo

提供给第三方ISV的的事件信息处理反馈的接口
*/
type AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest struct {
    model.Params
    // 系统上下文
    _param0   *WorkBenchContext
    // 请求封装类
    _param1   *EventInfoApiDto
}

// 初始化AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest对象
func NewAlibabaCampusDeviceOpenapiFeedbackeventinfoRequest() *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest{
    return &AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.feedbackeventinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统上下文
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 请求封装类
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) SetParam1(_param1 *EventInfoApiDto) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest) GetParam1() *EventInfoApiDto {
    return r._param1
}
