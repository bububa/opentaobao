package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeSimpleGetAPIResponse 获取交易订单的简易信息 API返回值
// taobao.trade.simple.get
//
// 获取CRM单表交易的详情
// &lt;br/&gt;1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// &lt;br/&gt;2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// &lt;br/&gt;3. 获取红包优惠金额可以使用字段 coupon_fee
// &lt;br/&gt;4. 请按需获取字段，减少TOP系统的压力
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoTradeSimpleGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeSimpleGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeSimpleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeSimpleGetAPIResponseModel).Reset()
}

// TaobaoTradeSimpleGetAPIResponseModel is 获取交易订单的简易信息 成功返回结果
type TaobaoTradeSimpleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_simple_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易主订单信息
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeSimpleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeSimpleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeSimpleGetAPIResponse)
	},
}

// GetTaobaoTradeSimpleGetAPIResponse 从 sync.Pool 获取 TaobaoTradeSimpleGetAPIResponse
func GetTaobaoTradeSimpleGetAPIResponse() *TaobaoTradeSimpleGetAPIResponse {
	return poolTaobaoTradeSimpleGetAPIResponse.Get().(*TaobaoTradeSimpleGetAPIResponse)
}

// ReleaseTaobaoTradeSimpleGetAPIResponse 将 TaobaoTradeSimpleGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeSimpleGetAPIResponse(v *TaobaoTradeSimpleGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeSimpleGetAPIResponse.Put(v)
}
