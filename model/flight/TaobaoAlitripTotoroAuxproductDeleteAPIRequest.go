package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTotoroAuxproductDeleteAPIRequest 廉航辅营产品删除 API请求
// taobao.alitrip.totoro.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaoAlitripTotoroAuxproductDeleteAPIRequest struct {
	model.Params
	// 廉航辅营产品删除请求
	_delAuxProductRq *DelAuxProductRq
}

// NewTaobaoAlitripTotoroAuxproductDeleteRequest 初始化TaobaoAlitripTotoroAuxproductDeleteAPIRequest对象
func NewTaobaoAlitripTotoroAuxproductDeleteRequest() *TaobaoAlitripTotoroAuxproductDeleteAPIRequest {
	return &TaobaoAlitripTotoroAuxproductDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTotoroAuxproductDeleteAPIRequest) Reset() {
	r._delAuxProductRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.totoro.auxproduct.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDelAuxProductRq is DelAuxProductRq Setter
// 廉航辅营产品删除请求
func (r *TaobaoAlitripTotoroAuxproductDeleteAPIRequest) SetDelAuxProductRq(_delAuxProductRq *DelAuxProductRq) error {
	r._delAuxProductRq = _delAuxProductRq
	r.Set("del_aux_product_rq", _delAuxProductRq)
	return nil
}

// GetDelAuxProductRq DelAuxProductRq Getter
func (r TaobaoAlitripTotoroAuxproductDeleteAPIRequest) GetDelAuxProductRq() *DelAuxProductRq {
	return r._delAuxProductRq
}

var poolTaobaoAlitripTotoroAuxproductDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTotoroAuxproductDeleteRequest()
	},
}

// GetTaobaoAlitripTotoroAuxproductDeleteRequest 从 sync.Pool 获取 TaobaoAlitripTotoroAuxproductDeleteAPIRequest
func GetTaobaoAlitripTotoroAuxproductDeleteAPIRequest() *TaobaoAlitripTotoroAuxproductDeleteAPIRequest {
	return poolTaobaoAlitripTotoroAuxproductDeleteAPIRequest.Get().(*TaobaoAlitripTotoroAuxproductDeleteAPIRequest)
}

// ReleaseTaobaoAlitripTotoroAuxproductDeleteAPIRequest 将 TaobaoAlitripTotoroAuxproductDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTotoroAuxproductDeleteAPIRequest(v *TaobaoAlitripTotoroAuxproductDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTotoroAuxproductDeleteAPIRequest.Put(v)
}
