package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务活动查询接口 APIResponse
taobao.jst.interactive.activity.query

互动任务活动查询接口
*/
type TaobaoJstInteractiveActivityQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveActivityQueryResponse
}

type TaobaoJstInteractiveActivityQueryResponse struct {
    XMLName xml.Name `xml:"jst_interactive_activity_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 活动列表，只返回有效的活动
    
    ActivityList   []Activity `json:"activity_list,omitempty" xml:"activity_list>activity,omitempty"`
    
    
}
