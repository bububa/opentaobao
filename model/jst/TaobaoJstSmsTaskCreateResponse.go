package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短信任务创建接口 APIResponse
taobao.jst.sms.task.create

聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
*/
type TaobaoJstSmsTaskCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_task_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 出参
    
    Result   *SmsResponse `json:"result,omitempty" xml:"