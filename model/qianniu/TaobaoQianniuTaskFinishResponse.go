package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
完成轻任务 APIResponse
taobao.qianniu.task.finish

由任务执行者调用
*/
type TaobaoQianniuTaskFinishAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskFinishResponse
}

type TaobaoQianniuTaskFinishResponse struct {
    XMLName xml.Name `xml:"qianniu_task_finish_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
