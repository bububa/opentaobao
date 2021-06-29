package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单同步 APIResponse
alitrip.btrip.city.car.apply.add

三方市内用车申请单同步
*/
type AlitripBtripCityCarApplyAddAPIResponse struct {
    model.CommonResponse
    AlitripBtripCityCarApplyAddResponse
}

type AlitripBtripCityCarApplyAddResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_city_car_apply_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 商旅内部审批单ID
    
    Module   int64 `json:"module,omitempty" xml:"module,omitempty"`

    
}
