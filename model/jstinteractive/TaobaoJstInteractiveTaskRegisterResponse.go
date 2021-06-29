package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务开通接口 APIResponse
taobao.jst.interactive.task.register

调用互动任务开通接口为小程序开通互动任务
*/
type TaobaoJstInteractiveTaskRegisterAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveTaskRegisterResponse
}

type TaobaoJstInteractiveTaskRegisterResponse struct {
    XMLName xml.Name `xml:"jst_interactive_task_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
