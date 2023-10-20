package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugOrderGetAPIResponse 查看订单详情 API返回值
// taobao.trade.drug.order.get
//
// 商家查看订单详情
type TaobaoTradeDrugOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeDrugOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeDrugOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeDrugOrderGetAPIResponseModel).Reset()
}

// TaobaoTradeDrugOrderGetAPIResponseModel is 查看订单详情 成功返回结果
type TaobaoTradeDrugOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoTradeDrugOrderGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeDrugOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTradeDrugOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeDrugOrderGetAPIResponse)
	},
}

// GetTaobaoTradeDrugOrderGetAPIResponse 从 sync.Pool 获取 TaobaoTradeDrugOrderGetAPIResponse
func GetTaobaoTradeDrugOrderGetAPIResponse() *TaobaoTradeDrugOrderGetAPIResponse {
	return poolTaobaoTradeDrugOrderGetAPIResponse.Get().(*TaobaoTradeDrugOrderGetAPIResponse)
}

// ReleaseTaobaoTradeDrugOrderGetAPIResponse 将 TaobaoTradeDrugOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeDrugOrderGetAPIResponse(v *TaobaoTradeDrugOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeDrugOrderGetAPIResponse.Put(v)
}
