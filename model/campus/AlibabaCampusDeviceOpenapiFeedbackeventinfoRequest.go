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
type AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest struct {
    model.Params
    // 系统上下文
    param0   *WorkBenchContext
    // 请求封装类
    param1   *EventInfoApiDto
}

// 初始化AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest对象
func NewAlibabaCampusDeviceOpenapiFeedbackeventinfoRequest() *AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest{
    return &AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.feedbackeventinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统上下文
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest) GetParam0() *WorkBenchContext {
    return r.param0
}
// Param1 Setter
// 请求封装类
func (r *AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest) SetParam1(param1 *EventInfoApiDto) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusDeviceOpenapiFeedbackeventinfoRequest) GetParam1() *EventInfoApiDto {
    return r.param1
}
