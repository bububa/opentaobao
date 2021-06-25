package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票保险推荐 APIRequest
taobao.alitrip.bus.tickets.insurance.recommend

获取推荐保险内容
*/
type TaobaoAlitripBusTicketsInsuranceRecommendRequest struct {
    model.Params

    // 请求对象
    recommendReq   *TopStandardInsRecommendRequest 

}

func NewTaobaoAlitripBusTicketsInsuranceRecommendRequest() *TaobaoAlitripBusTicketsInsuranceRecommendRequest{
    return &TaobaoAlitripBusTicketsInsuranceRecommendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripBusTicketsInsuranceRecommendRequest) GetApiMethodName() string {
    return "taobao.alitrip.bus.tickets.insurance.recommend"
}

func (r TaobaoAlitripBusTicketsInsuranceRecommendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripBusTicketsInsuranceRecommendRequest) SetRecommendReq(recommendReq *TopStandardInsRecommendRequest) error {
    r.recommendReq = recommendReq
    r.Set("recommend_req", recommendReq)
    return nil
}

func (r TaobaoAlitripBusTicketsInsuranceRecommendRequest) GetRecommendReq() *TopStandardInsRecommendRequest {
    return r.recommendReq
}

