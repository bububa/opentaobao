package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流订单呼叫运力 APIResponse
tmall.servicecenter.workcard.expressorder.consign

天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
*/
type TmallServicecenterWorkcardExpressorderConsignAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardExpressorderConsignResponse `json:"tmall_servicecenter_workcard_expressorder_consign_response,omitempty"`
}

type TmallServicecenterWorkcardExpressorderConsignResponse struct {

    // 返回结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
