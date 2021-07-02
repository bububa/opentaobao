package bus

import (
	"encoding/xml"

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

// TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel is 汽车票保险推荐 成功返回结果
type TaobaoAlitripBusTicketsInsuranceRecommendAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_bus_tickets_insurance_recommend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果数据
	Result *TaobaoAlitripBusTicketsInsuranceRecommendResult `json:"result,omitempty" xml:"result,omitempty"`
}
