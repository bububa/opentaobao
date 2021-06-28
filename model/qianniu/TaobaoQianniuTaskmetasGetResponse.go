package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务元查询接口 APIResponse
taobao.qianniu.taskmetas.get

任务元查询接口
*/
type TaobaoQianniuTaskmetasGetAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskmetasGetResponse
}

type TaobaoQianniuTaskmetasGetResponse struct {
    XMLName xml.Name `xml:"qianniu_taskmetas_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // taskmetas
    
    Taskmetas   []QTaskMetadata `json:"taskmetas,omitempty" xml:"taskmetas>q_task_metadata,omitempty"`
    
    
}
