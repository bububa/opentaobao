package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商制定维修费方案 API请求
cainiao.iot.ticket.sp.maintain.vtwo.create

服务商制定维修费方案
*/
type CainiaoIotTicketSpMaintainVtwoCreateRequest struct {
    model.Params
    // 维修方案
    makeMaintainPlanTopRequest   *MakeMaintainPlanV2TopRequest
}

// 初始化CainiaoIotTicketSpMaintainVtwoCreateRequest对象
func NewCainiaoIotTicketSpMaintainVtwoCreateRequest() *CainiaoIotTicketSpMaintainVtwoCreateRequest{
    return &CainiaoIotTicketSpMaintainVtwoCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpMaintainVtwoCreateRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.maintain.vtwo.create"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpMaintainVtwoCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MakeMaintainPlanTopRequest Setter
// 维修方案
func (r *CainiaoIotTicketSpMaintainVtwoCreateRequest) SetMakeMaintainPlanTopRequest(makeMaintainPlanTopRequest *MakeMaintainPlanV2TopRequest) error {
    r.makeMaintainPlanTopRequest = makeMaintainPlanTopRequest
    r.Set("make_maintain_plan_top_request", makeMaintainPlanTopRequest)
    return nil
}

// MakeMaintainPlanTopRequest Getter
func (r CainiaoIotTicketSpMaintainVtwoCreateRequest) GetMakeMaintainPlanTopRequest() *MakeMaintainPlanV2TopRequest {
    return r.makeMaintainPlanTopRequest
}
