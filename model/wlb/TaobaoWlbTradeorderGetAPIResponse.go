package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbTradeorderGetAPIResponse 根据交易号获取物流宝订单 API返回值
// taobao.wlb.tradeorder.get
//
// 根据交易类型和交易id查询物流宝订单详情
type TaobaoWlbTradeorderGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbTradeorderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbTradeorderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbTradeorderGetAPIResponseModel).Reset()
}

// TaobaoWlbTradeorderGetAPIResponseModel is 根据交易号获取物流宝订单 成功返回结果
type TaobaoWlbTradeorderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_tradeorder_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流宝订单对象
	WlbOrderList []WlbOrder `json:"wlb_order_list,omitempty" xml:"wlb_order_list>wlb_order,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbTradeorderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WlbOrderList = m.WlbOrderList[:0]
}

var poolTaobaoWlbTradeorderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbTradeorderGetAPIResponse)
	},
}

// GetTaobaoWlbTradeorderGetAPIResponse 从 sync.Pool 获取 TaobaoWlbTradeorderGetAPIResponse
func GetTaobaoWlbTradeorderGetAPIResponse() *TaobaoWlbTradeorderGetAPIResponse {
	return poolTaobaoWlbTradeorderGetAPIResponse.Get().(*TaobaoWlbTradeorderGetAPIResponse)
}

// ReleaseTaobaoWlbTradeorderGetAPIResponse 将 TaobaoWlbTradeorderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbTradeorderGetAPIResponse(v *TaobaoWlbTradeorderGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbTradeorderGetAPIResponse.Put(v)
}
