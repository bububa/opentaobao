package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse 汽车票保险推荐 API返回值
// taobao.alitrip.bus.tickets.insurance.recommend
//
// 获取推荐保险内容
type TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel).Reset()
}

// TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel is 汽车票保险推荐 成功返回结果
type TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bus_tickets_insurance_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果数据
	Result *TaobaoAlitripBusTicketsInsuranceRecommendResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse)
	},
}

// GetTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse 从 sync.Pool 获取 TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse
func GetTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse() *TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse {
	return poolTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse.Get().(*TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse)
}

// ReleaseTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse 将 TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse(v *TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse) {
	v.Reset()
	poolTaobaoAlitripBusTicketsInsuranceRecommendAPIResponse.Put(v)
}
