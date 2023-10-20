package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradesSoldQueryAPIResponse 根据收件人信息查询交易单号 API返回值
// taobao.trades.sold.query
//
// 根据收件人信息查询交易单号。
type TaobaoTradesSoldQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTradesSoldQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradesSoldQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradesSoldQueryAPIResponseModel).Reset()
}

// TaobaoTradesSoldQueryAPIResponseModel is 根据收件人信息查询交易单号 成功返回结果
type TaobaoTradesSoldQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"trades_sold_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单ID列表。按照订单创建时间倒序，最多返回最近的100笔订单。
	TidList []string `json:"tid_list,omitempty" xml:"tid_list>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradesSoldQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TidList = m.TidList[:0]
}

var poolTaobaoTradesSoldQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradesSoldQueryAPIResponse)
	},
}

// GetTaobaoTradesSoldQueryAPIResponse 从 sync.Pool 获取 TaobaoTradesSoldQueryAPIResponse
func GetTaobaoTradesSoldQueryAPIResponse() *TaobaoTradesSoldQueryAPIResponse {
	return poolTaobaoTradesSoldQueryAPIResponse.Get().(*TaobaoTradesSoldQueryAPIResponse)
}

// ReleaseTaobaoTradesSoldQueryAPIResponse 将 TaobaoTradesSoldQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradesSoldQueryAPIResponse(v *TaobaoTradesSoldQueryAPIResponse) {
	v.Reset()
	poolTaobaoTradesSoldQueryAPIResponse.Put(v)
}
