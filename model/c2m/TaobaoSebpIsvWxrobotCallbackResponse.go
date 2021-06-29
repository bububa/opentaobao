package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
isv机器人回调接口 APIResponse
taobao.sebp.isv.wxrobot.callback

机器人入群回调，进行校验、功能开通等操作
*/
type TaobaoSebpIsvWxrobotCallbackAPIResponse struct {
    model.CommonResponse
    TaobaoSebpIsvWxrobotCallbackResponse
}

type TaobaoSebpIsvWxrobotCallbackResponse struct {
    XMLName xml.Name `xml:"sebp_isv_wxrobot_callback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
