package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeMessagesGetAPIResponse 查询换货订单留言列表 API返回值
// tmall.exchange.messages.get
//
// 查询换货订单留言列表
type TmallExchangeMessagesGetAPIResponse struct {
	model.CommonResponse
	TmallExchangeMessagesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeMessagesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeMessagesGetAPIResponseModel).Reset()
}

// TmallExchangeMessagesGetAPIResponseModel is 查询换货订单留言列表 成功返回结果
type TmallExchangeMessagesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_messages_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RefundMessageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeMessagesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeMessagesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeMessagesGetAPIResponse)
	},
}

// GetTmallExchangeMessagesGetAPIResponse 从 sync.Pool 获取 TmallExchangeMessagesGetAPIResponse
func GetTmallExchangeMessagesGetAPIResponse() *TmallExchangeMessagesGetAPIResponse {
	return poolTmallExchangeMessagesGetAPIResponse.Get().(*TmallExchangeMessagesGetAPIResponse)
}

// ReleaseTmallExchangeMessagesGetAPIResponse 将 TmallExchangeMessagesGetAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeMessagesGetAPIResponse(v *TmallExchangeMessagesGetAPIResponse) {
	v.Reset()
	poolTmallExchangeMessagesGetAPIResponse.Put(v)
}
