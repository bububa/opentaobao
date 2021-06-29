package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务元查询接口 API返回值 
taobao.qianniu.taskmetas.get

任务元查询接口
*/
type TaobaoQianniuTaskmetasGetAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskmetasGetResponse
}

// 任务元查询接口 成功返回结果
type TaobaoQianniuTaskmetasGetResponse struct {
    XMLName xml.Name `xml:"qianniu_taskmetas_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // taskmetas
    Taskmetas   []QTaskMetadata `json:"taskmetas,omitempty" xml:"taskmetas>q_task_metadata,omitempty"`
}
