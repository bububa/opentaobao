package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品条码亮灯API APIResponse
taobao.uscesl.biz.item.light.up

亮灯API
*/
type TaobaoUsceslBizItemLightUpAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizItemLightUpResponse
}

type TaobaoUsceslBizItemLightUpResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_item_light_up_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoUsceslBizItemLightUpResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
