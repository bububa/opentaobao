package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
任务元查询接口 APIResponse
taobao.qianniu.taskmetas.get

任务元查询接口
*/
type TaobaoQianniuTaskmetasGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuTaskmetasGetResponse `json:"qianniu_taskmetas_get_response,omitempty"` 
    TaobaoQianniuTaskmetasGetResponse
}

/* model for simplify = false
type TaobaoQianniuTaskmetasGetResponse struct {

    // taskmetas
    
    Taskmetas  struct {
        QTaskMetadata  []QTaskMetadata `json:"q_task_metadata,omitempty"`
    } `json:"taskmetas,omitempty"`
    

}
*/

type TaobaoQianniuTaskmetasGetResponse struct {

    // taskmetas
    Taskmetas   []QTaskMetadata `json:"taskmetas,omitempty"`

}
