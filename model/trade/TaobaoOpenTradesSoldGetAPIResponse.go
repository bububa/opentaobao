package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenTradesSoldGetAPIResponse 查询卖家已卖出的交易数据（商家应用使用） API返回值
// taobao.open.trades.sold.get
//
// 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）&lt;br/&gt;
// 1. 返回的数据结果是以订单的创建时间倒序排列的。&lt;br/&gt;
// 注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。&lt;br/&gt;
// 2.入参fields中传入buyer_nick ，才能返回buyer_open_id
type TaobaoOpenTradesSoldGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenTradesSoldGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenTradesSoldGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenTradesSoldGetAPIResponseModel).Reset()
}

// TaobaoOpenTradesSoldGetAPIResponseModel is 查询卖家已卖出的交易数据（商家应用使用） 成功返回结果
type TaobaoOpenTradesSoldGetAPIResponseModel struct {
	XMLName xml.Name `xml:"open_trades_sold_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trades []Trade `json:"trades,omitempty" xml:"trades>trade,omitempty"`
	// 搜索到的交易信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenTradesSoldGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trades = m.Trades[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoOpenTradesSoldGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenTradesSoldGetAPIResponse)
	},
}

// GetTaobaoOpenTradesSoldGetAPIResponse 从 sync.Pool 获取 TaobaoOpenTradesSoldGetAPIResponse
func GetTaobaoOpenTradesSoldGetAPIResponse() *TaobaoOpenTradesSoldGetAPIResponse {
	return poolTaobaoOpenTradesSoldGetAPIResponse.Get().(*TaobaoOpenTradesSoldGetAPIResponse)
}

// ReleaseTaobaoOpenTradesSoldGetAPIResponse 将 TaobaoOpenTradesSoldGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenTradesSoldGetAPIResponse(v *TaobaoOpenTradesSoldGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenTradesSoldGetAPIResponse.Put(v)
}
