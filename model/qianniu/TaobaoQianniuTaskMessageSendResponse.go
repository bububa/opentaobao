package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发送任务提醒消息 APIResponse
taobao.qianniu.task.message.send

如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。
*/
type TaobaoQianniuTaskMessageSendAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQianniuTaskMessageSendResponse `json:"taobao_qianniu_task_message_send_response,omitempty"`
}

type TaobaoQianniuTaskMessageSendResponse struct {

    // 是否成功
    Result   bool `json:"result,omitempty"`

}
