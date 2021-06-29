package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单审批 APIResponse
alitrip.btrip.city.car.apply.approve

三方市内用车申请单审批
*/
type AlitripBtripCityCarApplyApproveAPIResponse struct {
    model.CommonResponse
    AlitripBtripCityCarApplyApproveResponse
}

type AlitripBtripCityCarApplyApproveResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_city_car_apply_approve_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 审批是否成功
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`

    
}
