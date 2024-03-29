package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeWtverticalGetAPIResponse 网厅垂直信息查询接口 API返回值
// taobao.trade.wtvertical.get
//
// 网厅订单垂直信息的查询
type TaobaoTradeWtverticalGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeWtverticalGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeWtverticalGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeWtverticalGetAPIResponseModel).Reset()
}

// TaobaoTradeWtverticalGetAPIResponseModel is 网厅垂直信息查询接口 成功返回结果
type TaobaoTradeWtverticalGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_wtvertical_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回交易垂直信息的数据结构
	WtextResults []WtExtResult `json:"wtext_results,omitempty" xml:"wtext_results>wt_ext_result,omitempty"`
	// 返回查询记录的条数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeWtverticalGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WtextResults = m.WtextResults[:0]
	m.TotalResults = 0
}

var poolTaobaoTradeWtverticalGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeWtverticalGetAPIResponse)
	},
}

// GetTaobaoTradeWtverticalGetAPIResponse 从 sync.Pool 获取 TaobaoTradeWtverticalGetAPIResponse
func GetTaobaoTradeWtverticalGetAPIResponse() *TaobaoTradeWtverticalGetAPIResponse {
	return poolTaobaoTradeWtverticalGetAPIResponse.Get().(*TaobaoTradeWtverticalGetAPIResponse)
}

// ReleaseTaobaoTradeWtverticalGetAPIResponse 将 TaobaoTradeWtverticalGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeWtverticalGetAPIResponse(v *TaobaoTradeWtverticalGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeWtverticalGetAPIResponse.Put(v)
}
