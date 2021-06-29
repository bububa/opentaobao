package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建轻任务 APIResponse
taobao.qianniu.task.create

发起一个轻任务，分配给多个执行者，并发送消息提醒，由任务发起者调用
*/
type TaobaoQianniuTaskCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskCreateResponse
}

type TaobaoQianniuTaskCreateResponse struct {
    XMLName xml.Name `xml:"qianniu_task_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 创建的任务元数据
    
    Result   *QTaskMetadata `json:"result,omitempty" xml:"result,omitempty"`

    
}
