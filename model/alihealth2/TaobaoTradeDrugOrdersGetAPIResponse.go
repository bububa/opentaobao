package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugOrdersGetAPIResponse 阿里健康获取某一药店全部订单 API返回值
// taobao.trade.drug.orders.get
//
// 阿里健康获取某一药店全部订单
type TaobaoTradeDrugOrdersGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeDrugOrdersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeDrugOrdersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeDrugOrdersGetAPIResponseModel).Reset()
}

// TaobaoTradeDrugOrdersGetAPIResponseModel is 阿里健康获取某一药店全部订单 成功返回结果
type TaobaoTradeDrugOrdersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_orders_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的订单列表对象
	Result *TaobaoTradeDrugOrdersGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeDrugOrdersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTradeDrugOrdersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeDrugOrdersGetAPIResponse)
	},
}

// GetTaobaoTradeDrugOrdersGetAPIResponse 从 sync.Pool 获取 TaobaoTradeDrugOrdersGetAPIResponse
func GetTaobaoTradeDrugOrdersGetAPIResponse() *TaobaoTradeDrugOrdersGetAPIResponse {
	return poolTaobaoTradeDrugOrdersGetAPIResponse.Get().(*TaobaoTradeDrugOrdersGetAPIResponse)
}

// ReleaseTaobaoTradeDrugOrdersGetAPIResponse 将 TaobaoTradeDrugOrdersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeDrugOrdersGetAPIResponse(v *TaobaoTradeDrugOrdersGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeDrugOrdersGetAPIResponse.Put(v)
}
