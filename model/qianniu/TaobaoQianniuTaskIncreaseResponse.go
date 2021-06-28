package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加任务接收人接口 APIResponse
taobao.qianniu.task.increase

根据任务元id增加任务接收人
*/
type TaobaoQianniuTaskIncreaseAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskIncreaseResponse
}

type TaobaoQianniuTaskIncreaseResponse struct {
    XMLName xml.Name `xml:"qianniu_task_increase_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否添加成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
