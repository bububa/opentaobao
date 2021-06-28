package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店线下已收尾款 APIResponse
tmall.car.fpcar.restpay.receive

提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
*/
type TmallCarFpcarRestpayReceiveAPIResponse struct {
    model.CommonResponse
    // Response *TmallCarFpcarRestpayReceiveResponse `json:"tmall_car_fpcar_restpay_receive_response,omitempty"` 
    TmallCarFpcarRestpayReceiveResponse
}

/* model for simplify = false
type TmallCarFpcarRestpayReceiveResponse struct {

    // 是否成功
    
    Succes   bool `json:"succes,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 业务结果信息详细描述
    
    Object   string `json:"object,omitempty"`
    

}
*/

type TmallCarFpcarRestpayReceiveResponse struct {

    // 是否成功
    Succes   bool `json:"succes,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 业务结果信息详细描述
    Object   string `json:"object,omitempty"`

}
