package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流订单创建API APIResponse
tmall.servicecenter.workcard.expressorder.create

天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
*/
type TmallServicecenterWorkcardExpressorderCreateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardExpressorderCreateResponse `json:"tmall_servicecenter_workcard_expressorder_create_response,omitempty"`
}

type TmallServicecenterWorkcardExpressorderCreateResponse struct {

    // 创建结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
