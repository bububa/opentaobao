package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分发放接口 APIResponse
taobao.jst.interactive.point.increase

向用户发放互动积分
*/
type TaobaoJstInteractivePointIncreaseAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractivePointIncreaseResponse
}

type TaobaoJstInteractivePointIncreaseResponse struct {
    XMLName xml.Name `xml:"jst_interactive_point_increase_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 用户积分总额
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`

    
}
