package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价签LED等点亮 APIResponse
taobao.uscesl.biz.light.up

价签LED等点亮
*/
type TaobaoUsceslBizLightUpAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizLightUpResponse
}

type TaobaoUsceslBizLightUpResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_light_up_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoUsceslBizLightUpResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
