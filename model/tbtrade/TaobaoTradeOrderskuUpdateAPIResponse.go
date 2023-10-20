package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeOrderskuUpdateAPIResponse 更新交易的销售属性 API返回值
// taobao.trade.ordersku.update
//
// 只能更新发货前子订单的销售属性 &lt;br/&gt;只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 &lt;br/&gt;必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
type TaobaoTradeOrderskuUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoTradeOrderskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeOrderskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeOrderskuUpdateAPIResponseModel).Reset()
}

// TaobaoTradeOrderskuUpdateAPIResponseModel is 更新交易的销售属性 成功返回结果
type TaobaoTradeOrderskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_ordersku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 只返回oid和modified
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeOrderskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Order = nil
}

var poolTaobaoTradeOrderskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeOrderskuUpdateAPIResponse)
	},
}

// GetTaobaoTradeOrderskuUpdateAPIResponse 从 sync.Pool 获取 TaobaoTradeOrderskuUpdateAPIResponse
func GetTaobaoTradeOrderskuUpdateAPIResponse() *TaobaoTradeOrderskuUpdateAPIResponse {
	return poolTaobaoTradeOrderskuUpdateAPIResponse.Get().(*TaobaoTradeOrderskuUpdateAPIResponse)
}

// ReleaseTaobaoTradeOrderskuUpdateAPIResponse 将 TaobaoTradeOrderskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeOrderskuUpdateAPIResponse(v *TaobaoTradeOrderskuUpdateAPIResponse) {
	v.Reset()
	poolTaobaoTradeOrderskuUpdateAPIResponse.Put(v)
}
