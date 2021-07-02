package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest 汽车票保险推荐 API请求
// taobao.alitrip.bus.tickets.insurance.recommend
//
// 获取推荐保险内容
type TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest struct {
	model.Params
	// 请求对象
	_recommendReq *TopStandardInsRecommendRequest
}

// NewTaobaoAlitripBusTicketsInsuranceRecommendRequest 初始化TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest对象
func NewTaobaoAlitripBusTicketsInsuranceRecommendRequest() *TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest {
	return &TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.bus.tickets.insurance.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RecommendReq Setter
// 请求对象
func (r *TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) SetRecommendReq(_recommendReq *TopStandardInsRecommendRequest) error {
	r._recommendReq = _recommendReq
	r.Set("recommend_req", _recommendReq)
	return nil
}

// Get RecommendReq Getter
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetRecommendReq() *TopStandardInsRecommendRequest {
	return r._recommendReq
}
