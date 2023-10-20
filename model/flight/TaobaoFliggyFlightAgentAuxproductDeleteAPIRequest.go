package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest 飞猪机票辅营商品删除 API请求
// taobao.fliggy.flight.agent.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest struct {
	model.Params
	// 廉航辅营产品删除请求
	_delAuxProductRq *DelAuxProductRq
}

// NewTaobaoFliggyFlightAgentAuxproductDeleteRequest 初始化TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest对象
func NewTaobaoFliggyFlightAgentAuxproductDeleteRequest() *TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest {
	return &TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) Reset() {
	r._delAuxProductRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fliggy.flight.agent.auxproduct.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDelAuxProductRq is DelAuxProductRq Setter
// 廉航辅营产品删除请求
func (r *TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) SetDelAuxProductRq(_delAuxProductRq *DelAuxProductRq) error {
	r._delAuxProductRq = _delAuxProductRq
	r.Set("del_aux_product_rq", _delAuxProductRq)
	return nil
}

// GetDelAuxProductRq DelAuxProductRq Getter
func (r TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) GetDelAuxProductRq() *DelAuxProductRq {
	return r._delAuxProductRq
}

var poolTaobaoFliggyFlightAgentAuxproductDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFliggyFlightAgentAuxproductDeleteRequest()
	},
}

// GetTaobaoFliggyFlightAgentAuxproductDeleteRequest 从 sync.Pool 获取 TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest
func GetTaobaoFliggyFlightAgentAuxproductDeleteAPIRequest() *TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest {
	return poolTaobaoFliggyFlightAgentAuxproductDeleteAPIRequest.Get().(*TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest)
}

// ReleaseTaobaoFliggyFlightAgentAuxproductDeleteAPIRequest 将 TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFliggyFlightAgentAuxproductDeleteAPIRequest(v *TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFliggyFlightAgentAuxproductDeleteAPIRequest.Put(v)
}
