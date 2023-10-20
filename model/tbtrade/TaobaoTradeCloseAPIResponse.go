package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeCloseAPIResponse 卖家关闭一笔交易 API返回值
// taobao.trade.close
//
// 关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
type TaobaoTradeCloseAPIResponse struct {
	model.CommonResponse
	TaobaoTradeCloseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeCloseAPIResponseModel).Reset()
}

// TaobaoTradeCloseAPIResponseModel is 卖家关闭一笔交易 成功返回结果
type TaobaoTradeCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关闭交易时返回的Trade信息，可用字段有tid和modified
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeCloseAPIResponse)
	},
}

// GetTaobaoTradeCloseAPIResponse 从 sync.Pool 获取 TaobaoTradeCloseAPIResponse
func GetTaobaoTradeCloseAPIResponse() *TaobaoTradeCloseAPIResponse {
	return poolTaobaoTradeCloseAPIResponse.Get().(*TaobaoTradeCloseAPIResponse)
}

// ReleaseTaobaoTradeCloseAPIResponse 将 TaobaoTradeCloseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeCloseAPIResponse(v *TaobaoTradeCloseAPIResponse) {
	v.Reset()
	poolTaobaoTradeCloseAPIResponse.Put(v)
}
