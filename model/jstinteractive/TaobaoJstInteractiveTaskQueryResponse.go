package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
互动任务列表查询接口 APIResponse
taobao.jst.interactive.task.query

查询用户的互动任务列表
*/
type TaobaoJstInteractiveTaskQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveTaskQueryResponse
}

type TaobaoJstInteractiveTaskQueryResponse struct {
    XMLName xml.Name `xml:"jst_interactive_task_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 查询结果
    
    Data   *InteractiveTaskQueryResponse `json:"data,omitempty" xml:"data,omitempty"`

    
}
