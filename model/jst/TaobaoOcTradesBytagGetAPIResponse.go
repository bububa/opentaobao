package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcTradesBytagGetAPIResponse 标签查询订单 API返回值
// taobao.oc.trades.bytag.get
//
// 根据标签查询订单编号
type TaobaoOcTradesBytagGetAPIResponse struct {
	model.CommonResponse
	TaobaoOcTradesBytagGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOcTradesBytagGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOcTradesBytagGetAPIResponseModel).Reset()
}

// TaobaoOcTradesBytagGetAPIResponseModel is 标签查询订单 成功返回结果
type TaobaoOcTradesBytagGetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_trades_bytag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 打了该标签的订单编号列表
	Tids []int64 `json:"tids,omitempty" xml:"tids>int64,omitempty"`
	// 总数
	Totals int64 `json:"totals,omitempty" xml:"totals,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOcTradesBytagGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Tids = m.Tids[:0]
	m.Totals = 0
}

var poolTaobaoOcTradesBytagGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOcTradesBytagGetAPIResponse)
	},
}

// GetTaobaoOcTradesBytagGetAPIResponse 从 sync.Pool 获取 TaobaoOcTradesBytagGetAPIResponse
func GetTaobaoOcTradesBytagGetAPIResponse() *TaobaoOcTradesBytagGetAPIResponse {
	return poolTaobaoOcTradesBytagGetAPIResponse.Get().(*TaobaoOcTradesBytagGetAPIResponse)
}

// ReleaseTaobaoOcTradesBytagGetAPIResponse 将 TaobaoOcTradesBytagGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOcTradesBytagGetAPIResponse(v *TaobaoOcTradesBytagGetAPIResponse) {
	v.Reset()
	poolTaobaoOcTradesBytagGetAPIResponse.Put(v)
}
