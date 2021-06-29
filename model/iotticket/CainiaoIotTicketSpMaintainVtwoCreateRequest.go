package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商制定维修费方案 APIRequest
cainiao.iot.ticket.sp.maintain.vtwo.create

服务商制定维修费方案
*/
type CainiaoIotTicketSpMaintainVtwoCreateRequest struct {
    model.Params

    // 维修方案
    makeMaintainPlanTopRequest   *MakeMaintainPlanV2TopRequest 

}

func NewCainiaoIotTicketSpMaintainVtwoCreateRequest() *CainiaoIotTicketSpMaintainVtwoCreateRequest{
    return &CainiaoIotTicketSpMaintainVtwoCreateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpMaintainVtwoCreateRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.maintain.vtwo.create"
}

func (r CainiaoIotTicketSpMaintainVtwoCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpMaintainVtwoCreateRequest) SetMakeMaintainPlanTopRequest(makeMaintainPlanTopRequest *MakeMaintainPlanV2TopRequest) error {
    r.makeMaintainPlanTopRequest = makeMaintainPlanTopRequest
    r.Set("make_maintain_plan_top_request", makeMaintainPlanTopRequest)
    return nil
}

func (r CainiaoIotTicketSpMaintainVtwoCreateRequest) GetMakeMaintainPlanTopRequest() *MakeMaintainPlanV2TopRequest {
    return r.makeMaintainPlanTopRequest
}

