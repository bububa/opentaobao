package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeMessageAddAPIResponse 卖家创建换货留言 API返回值
// tmall.exchange.message.add
//
// 卖家创建换货留言
type TmallExchangeMessageAddAPIResponse struct {
	model.CommonResponse
	TmallExchangeMessageAddAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeMessageAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeMessageAddAPIResponseModel).Reset()
}

// TmallExchangeMessageAddAPIResponseModel is 卖家创建换货留言 成功返回结果
type TmallExchangeMessageAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_message_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallExchangeMessageAddResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeMessageAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeMessageAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeMessageAddAPIResponse)
	},
}

// GetTmallExchangeMessageAddAPIResponse 从 sync.Pool 获取 TmallExchangeMessageAddAPIResponse
func GetTmallExchangeMessageAddAPIResponse() *TmallExchangeMessageAddAPIResponse {
	return poolTmallExchangeMessageAddAPIResponse.Get().(*TmallExchangeMessageAddAPIResponse)
}

// ReleaseTmallExchangeMessageAddAPIResponse 将 TmallExchangeMessageAddAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeMessageAddAPIResponse(v *TmallExchangeMessageAddAPIResponse) {
	v.Reset()
	poolTmallExchangeMessageAddAPIResponse.Put(v)
}
