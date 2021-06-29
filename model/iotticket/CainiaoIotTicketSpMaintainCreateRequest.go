package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商制定维修方案 APIRequest
cainiao.iot.ticket.sp.maintain.create

IoT售后服务商制定维修方案
*/
type CainiaoIotTicketSpMaintainCreateRequest struct {
    model.Params

    // 请求参数
    param   *AssignMaintenancePersonnelTopRequest 

}

func NewCainiaoIotTicketSpMaintainCreateRequest() *CainiaoIotTicketSpMaintainCreateRequest{
    return &CainiaoIotTicketSpMaintainCreateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpMaintainCreateRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.maintain.create"
}

func (r CainiaoIotTicketSpMaintainCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpMaintainCreateRequest) SetParam(param *AssignMaintenancePersonnelTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpMaintainCreateRequest) GetParam() *AssignMaintenancePersonnelTopRequest {
    return r.param
}

