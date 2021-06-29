package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动修改接口 APIResponse
taobao.jst.interactive.activity.update

互动任务活动修改接口
*/
type TaobaoJstInteractiveActivityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveActivityUpdateResponse
}

type TaobaoJstInteractiveActivityUpdateResponse struct {
    XMLName xml.Name `xml:"jst_interactive_activity_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 修改结果
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
