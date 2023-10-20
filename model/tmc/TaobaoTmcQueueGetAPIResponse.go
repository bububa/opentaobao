package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcqueuegetAPIResponse 获取消息队列积压情况 API返回值
// taobao.tmc.queue.get
//
// 根据appkey和groupName获取消息队列积压情况
type TaobaotmcqueuegetAPIResponse struct {
	model.CommonResponse
	TaobaotmcqueuegetAPIResponseModel
}

// TaobaotmcqueuegetAPIResponseModel is 获取消息队列积压情况 成功返回结果
type TaobaotmcqueuegetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_queue_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 队列详细信息
	Datas []TmcQueueInfo `json:"datas,omitempty" xml:"datas>tmc_queue_info,omitempty"`
}
