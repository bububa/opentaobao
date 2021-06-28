package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取消息队列积压情况 APIResponse
taobao.tmc.queue.get

根据appkey和groupName获取消息队列积压情况
*/
type TaobaoTmcQueueGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_queue_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 队列详细信息
    
    Datas   []TmcQueueInfo `json:"datas,omitempty" xml:"