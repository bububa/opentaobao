package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMessagesProduceAPIResponse 批量发送消息 API返回值
// taobao.tmc.messages.produce
//
// 批量发送消息
type TaobaoTmcMessagesProduceAPIResponse struct {
	model.CommonResponse
	TaobaoTmcMessagesProduceAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcMessagesProduceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcMessagesProduceAPIResponseModel).Reset()
}

// TaobaoTmcMessagesProduceAPIResponseModel is 批量发送消息 成功返回结果
type TaobaoTmcMessagesProduceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_messages_produce_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
	Results []TmcProduceResult `json:"results,omitempty" xml:"results>tmc_produce_result,omitempty"`
	// 是否全部成功
	IsAllSuccess bool `json:"is_all_success,omitempty" xml:"is_all_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcMessagesProduceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.IsAllSuccess = false
}

var poolTaobaoTmcMessagesProduceAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcMessagesProduceAPIResponse)
	},
}

// GetTaobaoTmcMessagesProduceAPIResponse 从 sync.Pool 获取 TaobaoTmcMessagesProduceAPIResponse
func GetTaobaoTmcMessagesProduceAPIResponse() *TaobaoTmcMessagesProduceAPIResponse {
	return poolTaobaoTmcMessagesProduceAPIResponse.Get().(*TaobaoTmcMessagesProduceAPIResponse)
}

// ReleaseTaobaoTmcMessagesProduceAPIResponse 将 TaobaoTmcMessagesProduceAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcMessagesProduceAPIResponse(v *TaobaoTmcMessagesProduceAPIResponse) {
	v.Reset()
	poolTaobaoTmcMessagesProduceAPIResponse.Put(v)
}
