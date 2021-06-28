package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
增加任务接收人接口 APIResponse
taobao.qianniu.task.increase

根据任务元id增加任务接收人
*/
type TaobaoQianniuTaskIncreaseAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuTaskIncreaseResponse `json:"qianniu_task_increase_response,omitempty"` 
    TaobaoQianniuTaskIncreaseResponse
}

/* model for simplify = false
type TaobaoQianniuTaskIncreaseResponse struct {

    // 是否添加成功
    
    Result   bool `json:"result,omitempty"`
    

}
*/

type TaobaoQianniuTaskIncreaseResponse struct {

    // 是否添加成功
    Result   bool `json:"result,omitempty"`

}
