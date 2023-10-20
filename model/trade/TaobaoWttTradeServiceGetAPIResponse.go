package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWttTradeServiceGetAPIResponse 获取网厅号卡垂直标信息 API返回值
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
type TaobaoWttTradeServiceGetAPIResponse struct {
	model.CommonResponse
	TaobaoWttTradeServiceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWttTradeServiceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWttTradeServiceGetAPIResponseModel).Reset()
}

// TaobaoWttTradeServiceGetAPIResponseModel is 获取网厅号卡垂直标信息 成功返回结果
type TaobaoWttTradeServiceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wtt_trade_service_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回实例
	Modules *WtverticalDto `json:"modules,omitempty" xml:"modules,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWttTradeServiceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modules = nil
}

var poolTaobaoWttTradeServiceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWttTradeServiceGetAPIResponse)
	},
}

// GetTaobaoWttTradeServiceGetAPIResponse 从 sync.Pool 获取 TaobaoWttTradeServiceGetAPIResponse
func GetTaobaoWttTradeServiceGetAPIResponse() *TaobaoWttTradeServiceGetAPIResponse {
	return poolTaobaoWttTradeServiceGetAPIResponse.Get().(*TaobaoWttTradeServiceGetAPIResponse)
}

// ReleaseTaobaoWttTradeServiceGetAPIResponse 将 TaobaoWttTradeServiceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWttTradeServiceGetAPIResponse(v *TaobaoWttTradeServiceGetAPIResponse) {
	v.Reset()
	poolTaobaoWttTradeServiceGetAPIResponse.Put(v)
}
