package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTotoroAuxproductDeleteAPIRequest
廉航辅营产品删除 API请求
taobao.alitrip.totoro.auxproduct.delete

廉航辅营产品删除接口 */
type TaobaoAlitripTotoroAuxproductDeleteAPIRequest struct {
	model.Params
	// 廉航辅营产品删除请求
	_delAuxProductRq *DelAuxProductRq
}

// NewTaobaoAlitripTotoroAuxproductDeleteRequest 初始化TaobaoAlitripTotoroAuxproductDeleteAPIRequest对象
func NewTaobaoAlitripTotoroAuxproductDeleteRequest() *TaobaoAlitripTotoroAuxproductDeleteAPIRequest {
	return &TaobaoAlitripTotoroAuxproductDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.totoro.auxproduct.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DelAuxProductRq Setter
// 廉航辅营产品删除请求
func (r *TaobaoAlitripTotoroAuxproductDeleteAPIRequest) SetDelAuxProductRq(_delAuxProductRq *DelAuxProductRq) error {
	r._delAuxProductRq = _delAuxProductRq
	r.Set("del_aux_product_rq", _delAuxProductRq)
	return nil
}

// Get DelAuxProductRq Getter
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetDelAuxProductRq() *DelAuxProductRq {
	return r._delAuxProductRq
}
