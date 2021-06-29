package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子价签解绑接口 APIResponse
taobao.uscesl.biz.esl.unbind

电子价签解绑接口
*/
type TaobaoUsceslBizEslUnbindAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizEslUnbindResponse
}

type TaobaoUsceslBizEslUnbindResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_esl_unbind_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result.sucess表示本次调用是否成功，true或false
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
