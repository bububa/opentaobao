package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店线下已收尾款 APIResponse
tmall.car.fpcar.restpay.receive

提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
*/
type TmallCarFpcarRestpayReceiveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_car_fpcar_restpay_receive_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Succes   bool `json:"succes,omitempty" xml:"