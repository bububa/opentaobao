package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取消息队列积压情况 APIResponse
taobao.tmc.queue.get

根据appkey和groupName获取消息队列积压情况
*/
type TaobaoTmcQueueGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcQueueGetResponse `json:"taobao_tmc_queue_get_response,omitempty"`
}

type TaobaoTmcQueueGetResponse struct {

    // 队列详细信息
    Datas   []TmcQueueInfo `json:"datas,omitempty"`

}
