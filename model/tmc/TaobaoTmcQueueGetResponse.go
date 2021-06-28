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
    TaobaoTmcQueueGetResponse
}

type TaobaoTmcQueueGetResponse struct {
    XMLName xml.Name `xml:"tmc_queue_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 队列详细信息
    
    Datas   []TmcQueueInfo `json:"datas,omitempty" xml:"datas>tmc_queue_info,omitempty"`
    
    
}
