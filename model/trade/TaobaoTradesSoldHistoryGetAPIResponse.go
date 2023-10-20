package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradesSoldHistoryGetAPIResponse 卖家历史库订单查询 API返回值
// taobao.trades.sold.history.get
//
// 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以上两年以内的交易信息）
// &lt;br/&gt;1. 返回的数据结果是以订单的创建时间倒序排列的。
// &lt;br/&gt;2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.simple.get获取订单信息。
// &lt;br/&gt;注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，解决办法就是type加上订单类型就可正常返回了。
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoTradesSoldHistoryGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradesSoldHistoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradesSoldHistoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradesSoldHistoryGetAPIResponseModel).Reset()
}

// TaobaoTradesSoldHistoryGetAPIResponseModel is 卖家历史库订单查询 成功返回结果
type TaobaoTradesSoldHistoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trades_sold_history_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trades []Trade `json:"trades,omitempty" xml:"trades>trade,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradesSoldHistoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trades = m.Trades[:0]
	m.HasNext = false
}

var poolTaobaoTradesSoldHistoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradesSoldHistoryGetAPIResponse)
	},
}

// GetTaobaoTradesSoldHistoryGetAPIResponse 从 sync.Pool 获取 TaobaoTradesSoldHistoryGetAPIResponse
func GetTaobaoTradesSoldHistoryGetAPIResponse() *TaobaoTradesSoldHistoryGetAPIResponse {
	return poolTaobaoTradesSoldHistoryGetAPIResponse.Get().(*TaobaoTradesSoldHistoryGetAPIResponse)
}

// ReleaseTaobaoTradesSoldHistoryGetAPIResponse 将 TaobaoTradesSoldHistoryGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradesSoldHistoryGetAPIResponse(v *TaobaoTradesSoldHistoryGetAPIResponse) {
	v.Reset()
	poolTaobaoTradesSoldHistoryGetAPIResponse.Put(v)
}
