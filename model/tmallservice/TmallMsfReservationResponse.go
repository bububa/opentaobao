package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅服务预约API APIResponse
tmall.msf.reservation

喵师傅预约api
*/
type TmallMsfReservationAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_msf_reservation_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 预约成功,json
    
    Result   string `json:"result,omitempty" xml:"