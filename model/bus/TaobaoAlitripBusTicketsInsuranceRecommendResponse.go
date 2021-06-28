package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票保险推荐 APIResponse
taobao.alitrip.bus.tickets.insurance.recommend

获取推荐保险内容
*/
type TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripBusTicketsInsuranceRecommendResponse
}

type TaobaoAlitripBusTicketsInsuranceRecommendResponse struct {
    XMLName xml.Name `xml:"alitrip_bus_tickets_insurance_recommend_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回结果数据
    
    Result   *TaobaoAlitripBusTicketsInsuranceRecommendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
