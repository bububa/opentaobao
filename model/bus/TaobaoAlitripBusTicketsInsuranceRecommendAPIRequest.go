package bus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) Reset() {
	r._recommendReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.bus.tickets.insurance.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecommendReq is RecommendReq Setter
// 请求对象
func (r *TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) SetRecommendReq(_recommendReq *TopStandardInsRecommendRequest) error {
	r._recommendReq = _recommendReq
	r.Set("recommend_req", _recommendReq)
	return nil
}

// GetRecommendReq RecommendReq Getter
func (r TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) GetRecommendReq() *TopStandardInsRecommendRequest {
	return r._recommendReq
}

var poolTaobaoAlitripBusTicketsInsuranceRecommendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripBusTicketsInsuranceRecommendRequest()
	},
}

// GetTaobaoAlitripBusTicketsInsuranceRecommendRequest 从 sync.Pool 获取 TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest
func GetTaobaoAlitripBusTicketsInsuranceRecommendAPIRequest() *TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest {
	return poolTaobaoAlitripBusTicketsInsuranceRecommendAPIRequest.Get().(*TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest)
}

// ReleaseTaobaoAlitripBusTicketsInsuranceRecommendAPIRequest 将 TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripBusTicketsInsuranceRecommendAPIRequest(v *TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest) {
	v.Reset()
	poolTaobaoAlitripBusTicketsInsuranceRecommendAPIRequest.Put(v)
}
