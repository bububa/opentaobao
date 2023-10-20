package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductskusgetAPIRequest SKU查询接口 API请求
// taobao.fenxiao.product.skus.get
//
// 产品sku查询
type TaobaofenxiaoproductskusgetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// NewTaobaofenxiaoproductskusgetRequest 初始化TaobaofenxiaoproductskusgetAPIRequest对象
func NewTaobaofenxiaoproductskusgetRequest() *TaobaofenxiaoproductskusgetAPIRequest {
	return &TaobaofenxiaoproductskusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductskusgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.skus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductskusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductskusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaofenxiaoproductskusgetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductskusgetAPIRequest) GetProductId() int64 {
	return r._productId
}
