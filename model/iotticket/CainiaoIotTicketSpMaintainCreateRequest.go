package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商制定维修方案 API请求
cainiao.iot.ticket.sp.maintain.create

IoT售后服务商制定维修方案
*/
type CainiaoIotTicketSpMaintainCreateRequest struct {
    model.Params
    // 请求参数
    param   *AssignMaintenancePersonnelTopRequest
}

// 初始化CainiaoIotTicketSpMaintainCreateRequest对象
func NewCainiaoIotTicketSpMaintainCreateRequest() *CainiaoIotTicketSpMaintainCreateRequest{
    return &CainiaoIotTicketSpMaintainCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainCreateRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.maintain.create"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpMaintainCreateRequest) SetParam(param *AssignMaintenancePersonnelTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpMaintainCreateRequest) GetParam() *AssignMaintenancePersonnelTopRequest {
    return r.param
}
