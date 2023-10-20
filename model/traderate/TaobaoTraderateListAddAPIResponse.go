package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTraderateListAddAPIResponse 针对父子订单新增批量评价 API返回值
// taobao.traderate.list.add
//
// 针对父子订单新增批量评价(&lt;font color=&#34;red&#34;&gt;注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价&lt;/font&gt;)
type TaobaoTraderateListAddAPIResponse struct {
	model.CommonResponse
	TaobaoTraderateListAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTraderateListAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTraderateListAddAPIResponseModel).Reset()
}

// TaobaoTraderateListAddAPIResponseModel is 针对父子订单新增批量评价 成功返回结果
type TaobaoTraderateListAddAPIResponseModel struct {
	XMLName xml.Name `xml:"traderate_list_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的评论的信息，仅返回tid和created字段
	TradeRate *TradeRateRequest `json:"trade_rate,omitempty" xml:"trade_rate,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTraderateListAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeRate = nil
}

var poolTaobaoTraderateListAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTraderateListAddAPIResponse)
	},
}

// GetTaobaoTraderateListAddAPIResponse 从 sync.Pool 获取 TaobaoTraderateListAddAPIResponse
func GetTaobaoTraderateListAddAPIResponse() *TaobaoTraderateListAddAPIResponse {
	return poolTaobaoTraderateListAddAPIResponse.Get().(*TaobaoTraderateListAddAPIResponse)
}

// ReleaseTaobaoTraderateListAddAPIResponse 将 TaobaoTraderateListAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTraderateListAddAPIResponse(v *TaobaoTraderateListAddAPIResponse) {
	v.Reset()
	poolTaobaoTraderateListAddAPIResponse.Put(v)
}
