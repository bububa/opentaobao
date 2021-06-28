package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单预约 APIResponse
tmall.servicecenter.workcard.reserve

服务工单更新通用接口
*/
type TmallServicecenterWorkcardReserveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_reserve_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *TmallServicecenterWorkcardReserveResult `json:"result,omitempty" xml:"