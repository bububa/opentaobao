package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptotoroauxproductdeleteAPIRequest 廉航辅营产品删除 API请求
// taobao.alitrip.totoro.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaoalitriptotoroauxproductdeleteAPIRequest struct {
	model.Params
	// 廉航辅营产品删除请求
	_delAuxProductRq *DelAuxProductRq
}

// NewTaobaoalitriptotoroauxproductdeleteRequest 初始化TaobaoalitriptotoroauxproductdeleteAPIRequest对象
func NewTaobaoalitriptotoroauxproductdeleteRequest() *TaobaoalitriptotoroauxproductdeleteAPIRequest {
	return &TaobaoalitriptotoroauxproductdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptotoroauxproductdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.totoro.auxproduct.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptotoroauxproductdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptotoroauxproductdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDelAuxProductRq is DelAuxProductRq Setter
// 廉航辅营产品删除请求
func (r *TaobaoalitriptotoroauxproductdeleteAPIRequest) SetDelAuxProductRq(_delAuxProductRq *DelAuxProductRq) error {
	r._delAuxProductRq = _delAuxProductRq
	r.Set("del_aux_product_rq", _delAuxProductRq)
	return nil
}

// GetDelAuxProductRq DelAuxProductRq Getter
func (r TaobaoalitriptotoroauxproductdeleteAPIRequest) GetDelAuxProductRq() *DelAuxProductRq {
	return r._delAuxProductRq
}
