package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票保险推荐 APIResponse
taobao.alitrip.bus.tickets.insurance.recommend

获取推荐保险内容
*/
type TaobaoAlitripBusTicketsInsuranceRecommendAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripBusTicketsInsuranceRecommendResponse `json:"alitrip_bus_tickets_insurance_recommend_response,omitempty"` 
    TaobaoAlitripBusTicketsInsuranceRecommendResponse
}

/* model for simplify = false
type TaobaoAlitripBusTicketsInsuranceRecommendResponse struct {

    // 接口返回结果数据
    
    Result  *struct {
        TaobaoAlitripBusTicketsInsuranceRecommendResult  *TaobaoAlitripBusTicketsInsuranceRecommendResult `json:"taobao_alitrip_bus_tickets_insurance_recommend_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAlitripBusTicketsInsuranceRecommendResponse struct {

    // 接口返回结果数据
    Result   *TaobaoAlitripBusTicketsInsuranceRecommendResult `json:"result,omitempty"`

}
