package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
轻任务删除接口 APIResponse
taobao.qianniu.task.remove

轻任务删除接口。
*/
type TaobaoQianniuTaskRemoveAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskRemoveResponse
}

type TaobaoQianniuTaskRemoveResponse struct {
    XMLName xml.Name `xml:"qianniu_task_remove_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
