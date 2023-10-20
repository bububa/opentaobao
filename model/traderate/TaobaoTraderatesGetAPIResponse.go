package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTraderatesGetAPIResponse 搜索评价信息 API返回值
// taobao.traderates.get
//
// 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
type TaobaoTraderatesGetAPIResponse struct {
	model.CommonResponse
	TaobaoTraderatesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTraderatesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTraderatesGetAPIResponseModel).Reset()
}

// TaobaoTraderatesGetAPIResponseModel is 搜索评价信息 成功返回结果
type TaobaoTraderatesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"traderates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息
	TradeRates []TaobaoTraderatesGetResults `json:"trade_rates,omitempty" xml:"trade_rates>taobao_traderates_get_results,omitempty"`
	// 搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 当使用use_has_next时返回信息，如果还有下一页则返回true
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTraderatesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeRates = m.TradeRates[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoTraderatesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTraderatesGetAPIResponse)
	},
}

// GetTaobaoTraderatesGetAPIResponse 从 sync.Pool 获取 TaobaoTraderatesGetAPIResponse
func GetTaobaoTraderatesGetAPIResponse() *TaobaoTraderatesGetAPIResponse {
	return poolTaobaoTraderatesGetAPIResponse.Get().(*TaobaoTraderatesGetAPIResponse)
}

// ReleaseTaobaoTraderatesGetAPIResponse 将 TaobaoTraderatesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTraderatesGetAPIResponse(v *TaobaoTraderatesGetAPIResponse) {
	v.Reset()
	poolTaobaoTraderatesGetAPIResponse.Put(v)
}
