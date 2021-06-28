package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店通知用户提车 APIResponse
tmall.car.fpcar.getcar.notify

提供给外部(大搜或其它合作方)的接口-门店通知用户提车
*/
type TmallCarFpcarGetcarNotifyAPIResponse struct {
    model.CommonResponse
    // Response *TmallCarFpcarGetcarNotifyResponse `json:"tmall_car_fpcar_getcar_notify_response,omitempty"` 
    TmallCarFpcarGetcarNotifyResponse
}

/* model for simplify = false
type TmallCarFpcarGetcarNotifyResponse struct {

    // 返回的数据结果
    
    Object   string `json:"object,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 是否成功
    
    Succes   bool `json:"succes,omitempty"`
    

}
*/

type TmallCarFpcarGetcarNotifyResponse struct {

    // 返回的数据结果
    Object   string `json:"object,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // 是否成功
    Succes   bool `json:"succes,omitempty"`

}
