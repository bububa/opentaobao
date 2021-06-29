package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商维修方案更新 API请求
cainiao.iot.ticket.sp.maintain.update

IoT售后服务商维修方案更新
*/
type CainiaoIotTicketSpMaintainUpdateRequest struct {
    model.Params
    // 请求参数
    param   *UpdateMaintainPlanTopRequest
}

// 初始化CainiaoIotTicketSpMaintainUpdateRequest对象
func NewCainiaoIotTicketSpMaintainUpdateRequest() *CainiaoIotTicketSpMaintainUpdateRequest{
    return &CainiaoIotTicketSpMaintainUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainUpdateRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.maintain.update"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMaintainUpdateRequest) SetParam(param *UpdateMaintainPlanTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpMaintainUpdateRequest) GetParam() *UpdateMaintainPlanTopRequest {
    return r.param
}
