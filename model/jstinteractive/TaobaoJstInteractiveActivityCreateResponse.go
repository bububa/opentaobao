package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动创建接口 APIResponse
taobao.jst.interactive.activity.create

调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回
*/
type TaobaoJstInteractiveActivityCreateAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveActivityCreateResponse
}

type TaobaoJstInteractiveActivityCreateResponse struct {
    XMLName xml.Name `xml:"jst_interactive_activity_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 活动id
    
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`

    
}
