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
    TmallServicecenterWorkcardReservefailResponse
}

type TmallServicecenterWorkcardReservefailResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_reservefail_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // -
    
    Result   *TmallServicecenterWorkcardReservefailResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
