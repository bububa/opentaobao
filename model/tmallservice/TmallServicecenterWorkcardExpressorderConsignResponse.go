package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单呼叫运力 APIResponse
tmall.servicecenter.workcard.expressorder.consign

天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
*/
type TmallServicecenterWorkcardExpressorderConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_expressorder_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"