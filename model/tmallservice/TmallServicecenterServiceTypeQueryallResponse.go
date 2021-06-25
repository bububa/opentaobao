package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务类型 APIResponse
tmall.servicecenter.service.type.queryall

查询天猫服务类型列表
*/
type TmallServicecenterServiceTypeQueryallAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServiceTypeQueryallResponse `json:"tmall_servicecenter_service_type_queryall_response,omitempty"`
}

type TmallServicecenterServiceTypeQueryallResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}