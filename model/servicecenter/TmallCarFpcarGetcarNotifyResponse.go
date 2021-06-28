package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店通知用户提车 APIResponse
tmall.car.fpcar.getcar.notify

提供给外部(大搜或其它合作方)的接口-门店通知用户提车
*/
type TmallCarFpcarGetcarNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_car_fpcar_getcar_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的数据结果
    
    Object   string `json:"object,omitempty" xml:"