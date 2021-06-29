package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商维修方案更新 APIRequest
cainiao.iot.ticket.sp.maintain.update

IoT售后服务商维修方案更新
*/
type CainiaoIotTicketSpMaintainUpdateRequest struct {
    model.Params

    // 请求参数
    param   *UpdateMaintainPlanTopRequest 

}

func NewCainiaoIotTicketSpMaintainUpdateRequest() *CainiaoIotTicketSpMaintainUpdateRequest{
    return &CainiaoIotTicketSpMaintainUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpMaintainUpdateRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.maintain.update"
}

func (r CainiaoIotTicketSpMaintainUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpMaintainUpdateRequest) SetParam(param *UpdateMaintainPlanTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpMaintainUpdateRequest) GetParam() *UpdateMaintainPlanTopRequest {
    return r.param
}

