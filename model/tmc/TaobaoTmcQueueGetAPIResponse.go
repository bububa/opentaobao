package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcQueueGetAPIResponse 获取消息队列积压情况 API返回值
// taobao.tmc.queue.get
//
// 根据appkey和groupName获取消息队列积压情况
type TaobaoTmcQueueGetAPIResponse struct {
	model.CommonResponse
	TaobaoTmcQueueGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcQueueGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcQueueGetAPIResponseModel).Reset()
}

// TaobaoTmcQueueGetAPIResponseModel is 获取消息队列积压情况 成功返回结果
type TaobaoTmcQueueGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_queue_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 队列详细信息
	Datas []TmcQueueInfo `json:"datas,omitempty" xml:"datas>tmc_queue_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcQueueGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
}

var poolTaobaoTmcQueueGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcQueueGetAPIResponse)
	},
}

// GetTaobaoTmcQueueGetAPIResponse 从 sync.Pool 获取 TaobaoTmcQueueGetAPIResponse
func GetTaobaoTmcQueueGetAPIResponse() *TaobaoTmcQueueGetAPIResponse {
	return poolTaobaoTmcQueueGetAPIResponse.Get().(*TaobaoTmcQueueGetAPIResponse)
}

// ReleaseTaobaoTmcQueueGetAPIResponse 将 TaobaoTmcQueueGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcQueueGetAPIResponse(v *TaobaoTmcQueueGetAPIResponse) {
	v.Reset()
	poolTaobaoTmcQueueGetAPIResponse.Put(v)
}
