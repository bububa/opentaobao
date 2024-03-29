package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeFullinfoGetAPIResponse 获取单笔交易的详细信息 API返回值
// taobao.trade.fullinfo.get
//
// 获取单笔交易的详细信息
// &lt;br/&gt;1. 只有单笔订单的情况下Trade数据结构中才包含商品相关的信息
// &lt;br/&gt;2. 获取到的Order中的payment字段在单笔子订单时包含物流费用，多笔子订单时不包含物流费用
// &lt;br/&gt;3. 获取红包金额使用字段：tmall_coupon_fee（天猫商家订单使用的红包金额）
// &lt;br/&gt;4. 请按需获取字段，减少TOP系统的压力
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=tradeapi&#34; target=&#34;_blank&#34;&gt;点击查看更多交易API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoTradeFullinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeFullinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeFullinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeFullinfoGetAPIResponseModel).Reset()
}

// TaobaoTradeFullinfoGetAPIResponseModel is 获取单笔交易的详细信息 成功返回结果
type TaobaoTradeFullinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_fullinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易主订单信息
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeFullinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeFullinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeFullinfoGetAPIResponse)
	},
}

// GetTaobaoTradeFullinfoGetAPIResponse 从 sync.Pool 获取 TaobaoTradeFullinfoGetAPIResponse
func GetTaobaoTradeFullinfoGetAPIResponse() *TaobaoTradeFullinfoGetAPIResponse {
	return poolTaobaoTradeFullinfoGetAPIResponse.Get().(*TaobaoTradeFullinfoGetAPIResponse)
}

// ReleaseTaobaoTradeFullinfoGetAPIResponse 将 TaobaoTradeFullinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeFullinfoGetAPIResponse(v *TaobaoTradeFullinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeFullinfoGetAPIResponse.Put(v)
}
