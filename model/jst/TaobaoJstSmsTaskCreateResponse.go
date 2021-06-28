package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短信任务创建接口 APIResponse
taobao.jst.sms.task.create

聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
*/
type TaobaoJstSmsTaskCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstSmsTaskCreateResponse `json:"jst_sms_task_create_response,omitempty"` 
    TaobaoJstSmsTaskCreateResponse
}

/* model for simplify = false
type TaobaoJstSmsTaskCreateResponse struct {

    // 出参
    
    Result  *struct {
        SmsResponse  *SmsResponse `json:"sms_response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoJstSmsTaskCreateResponse struct {

    // 出参
    Result   *SmsResponse `json:"result,omitempty"`

}
