package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofliggyflightagentauxproductdeleteAPIRequest 飞猪机票辅营商品删除 API请求
// taobao.fliggy.flight.agent.auxproduct.delete
//
// 廉航辅营产品删除接口
type TaobaofliggyflightagentauxproductdeleteAPIRequest struct {
	model.Params
	// 廉航辅营产品删除请求
	_delAuxProductRq *DelAuxProductRq
}

// NewTaobaofliggyflightagentauxproductdeleteRequest 初始化TaobaofliggyflightagentauxproductdeleteAPIRequest对象
func NewTaobaofliggyflightagentauxproductdeleteRequest() *TaobaofliggyflightagentauxproductdeleteAPIRequest {
	return &TaobaofliggyflightagentauxproductdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofliggyflightagentauxproductdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fliggy.flight.agent.auxproduct.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofliggyflightagentauxproductdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofliggyflightagentauxproductdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDelAuxProductRq is DelAuxProductRq Setter
// 廉航辅营产品删除请求
func (r *TaobaofliggyflightagentauxproductdeleteAPIRequest) SetDelAuxProductRq(_delAuxProductRq *DelAuxProductRq) error {
	r._delAuxProductRq = _delAuxProductRq
	r.Set("del_aux_product_rq", _delAuxProductRq)
	return nil
}

// GetDelAuxProductRq DelAuxProductRq Getter
func (r TaobaofliggyflightagentauxproductdeleteAPIRequest) GetDelAuxProductRq() *DelAuxProductRq {
	return r._delAuxProductRq
}
