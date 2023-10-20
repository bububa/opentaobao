package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugGetAPIResponse 查询商家未确认订单列表 API返回值
// taobao.trade.drug.get
//
// 可以按商家或是店铺维度的进行查询买家付款卖家未确认订单，一次返回不大于20条订单
type TaobaoTradeDrugGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeDrugGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeDrugGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeDrugGetAPIResponseModel).Reset()
}

// TaobaoTradeDrugGetAPIResponseModel is 查询商家未确认订单列表 成功返回结果
type TaobaoTradeDrugGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_drug_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询到的订单列表对象
	Result *TaobaoTradeDrugGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeDrugGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTradeDrugGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeDrugGetAPIResponse)
	},
}

// GetTaobaoTradeDrugGetAPIResponse 从 sync.Pool 获取 TaobaoTradeDrugGetAPIResponse
func GetTaobaoTradeDrugGetAPIResponse() *TaobaoTradeDrugGetAPIResponse {
	return poolTaobaoTradeDrugGetAPIResponse.Get().(*TaobaoTradeDrugGetAPIResponse)
}

// ReleaseTaobaoTradeDrugGetAPIResponse 将 TaobaoTradeDrugGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeDrugGetAPIResponse(v *TaobaoTradeDrugGetAPIResponse) {
	v.Reset()
	poolTaobaoTradeDrugGetAPIResponse.Put(v)
}
