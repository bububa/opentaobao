package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单申请接口 APIResponse
taobao.fuwu.sp.confirm.apply

isv能通过该接口发起确认申请单
*/
type TaobaoFuwuSpConfirmApplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fuwu_sp_confirm_apply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的是服务市场的确认单ID
    
    ApplyResult   int64 `json:"apply_result,omitempty" xml:"