package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票保险推荐 API请求
taobao.alitrip.bus.tickets.insurance.recommend

获取推荐保险内容
*/
type TaobaoAlitripBusTicketsInsuranceRecommendRequest struct {
    model.Params
    // 请求对象
    _recommendReq   *TopStandardInsRecommendRequest
}

// 初始化TaobaoAlitripBusTicketsInsuranceRecommendRequest对象
func NewTaobaoAlitripBusTicketsInsuranceRecommendRequest() *TaobaoAlitripBusTicketsInsuranceRecommendRequest{
    return &TaobaoAlitripBusTicketsInsuranceRecommendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripBusTicketsInsuranceRecommendRequest) GetApiMethodName() string {
    return "taobao.alitrip.bus.tickets.insurance.recommend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripBusTicketsInsuranceRecommendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RecommendReq Setter
// 请求对象
func (r *TaobaoAlitripBusTicketsInsuranceRecommendRequest) SetRecommendReq(_recommendReq *TopStandardInsRecommendRequest) error {
    r._recommendReq = _recommendReq
    r.Set("recommend_req", _recommendReq)
    return nil
}

// RecommendReq Getter
func (r TaobaoAlitripBusTicketsInsuranceRecommendRequest) GetRecommendReq() *TopStandardInsRecommendRequest {
    return r._recommendReq
}
