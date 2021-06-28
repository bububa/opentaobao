package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预约失败 APIResponse
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败
*/
type TmallServicecenterWorkcardReservefailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_reservefail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // -
    
    Result   *TmallServicecenterWorkcardReservefailResult `json:"result,omitempty" xml:"