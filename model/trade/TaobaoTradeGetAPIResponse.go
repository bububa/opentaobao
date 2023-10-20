package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeGetAPIResponse 获取单笔交易的部分信息(性能高) API返回值
// taobao.trade.get
//
// 获取单笔交易的部分信息
// &lt;br/&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;
type TaobaoTradeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeGetAPIResponseModel).Reset()
}

// TaobaoTradeGetAPIResponseModel is 获取单笔交易的部分信息(性能高) 成功返回结果
type TaobaoTradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeGetAPIResponse)
	},
}

// GetTaobaoTradeGetAPIResponse 从 sync.Pool 获取 TaobaoTradeGetAPIResponse
func GetTaobaoTradeGetAPIResponse() *TaobaoTradeGetAPIResponse {
	return poolTaobaoTradeGetAPIResponse.Get().(*TaobaoTradeGetAPIResponse)
}

// ReleaseTaobaoTradeGetAPIResponse 将 TaobaoTradeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeGetAPIResponse(v *TaobaoTradeGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeGetAPIResponse.Put(v)
}
