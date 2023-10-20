package aliexpresssumaitong

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeOrderOpenQueryAPIResponse Aliexpress开放平台订单查询 API返回值
// aliexpress.trade.order.open.query
//
// Aliexpress开放平台订单信息查询
type AliexpressTradeOrderOpenQueryAPIResponse struct {
	model.CommonResponse
	AliexpressTradeOrderOpenQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressTradeOrderOpenQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressTradeOrderOpenQueryAPIResponseModel).Reset()
}

// AliexpressTradeOrderOpenQueryAPIResponseModel is Aliexpress开放平台订单查询 成功返回结果
type AliexpressTradeOrderOpenQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_trade_order_open_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单查询接口返回值
	Result *OrderQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressTradeOrderOpenQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressTradeOrderOpenQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressTradeOrderOpenQueryAPIResponse)
	},
}

// GetAliexpressTradeOrderOpenQueryAPIResponse 从 sync.Pool 获取 AliexpressTradeOrderOpenQueryAPIResponse
func GetAliexpressTradeOrderOpenQueryAPIResponse() *AliexpressTradeOrderOpenQueryAPIResponse {
	return poolAliexpressTradeOrderOpenQueryAPIResponse.Get().(*AliexpressTradeOrderOpenQueryAPIResponse)
}

// ReleaseAliexpressTradeOrderOpenQueryAPIResponse 将 AliexpressTradeOrderOpenQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressTradeOrderOpenQueryAPIResponse(v *AliexpressTradeOrderOpenQueryAPIResponse) {
	v.Reset()
	poolAliexpressTradeOrderOpenQueryAPIResponse.Put(v)
}
