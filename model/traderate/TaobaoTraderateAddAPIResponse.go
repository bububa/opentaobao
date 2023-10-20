package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTraderateAddAPIResponse 新增单个评价 API返回值
// taobao.traderate.add
//
// 新增单个评价(&lt;font color=&#34;red&#34;&gt;注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价&lt;/font&gt;)
type TaobaoTraderateAddAPIResponse struct {
	model.CommonResponse
	TaobaoTraderateAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTraderateAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTraderateAddAPIResponseModel).Reset()
}

// TaobaoTraderateAddAPIResponseModel is 新增单个评价 成功返回结果
type TaobaoTraderateAddAPIResponseModel struct {
	XMLName xml.Name `xml:"traderate_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回tid、oid、create
	TradeRate *TradeRateRequest `json:"trade_rate,omitempty" xml:"trade_rate,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTraderateAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeRate = nil
}

var poolTaobaoTraderateAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTraderateAddAPIResponse)
	},
}

// GetTaobaoTraderateAddAPIResponse 从 sync.Pool 获取 TaobaoTraderateAddAPIResponse
func GetTaobaoTraderateAddAPIResponse() *TaobaoTraderateAddAPIResponse {
	return poolTaobaoTraderateAddAPIResponse.Get().(*TaobaoTraderateAddAPIResponse)
}

// ReleaseTaobaoTraderateAddAPIResponse 将 TaobaoTraderateAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTraderateAddAPIResponse(v *TaobaoTraderateAddAPIResponse) {
	v.Reset()
	poolTaobaoTraderateAddAPIResponse.Put(v)
}
