package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkusGetAPIRequest SKU查询接口 API请求
// taobao.fenxiao.product.skus.get
//
// 产品sku查询
type TaobaoFenxiaoProductSkusGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
}

// NewTaobaoFenxiaoProductSkusGetRequest 初始化TaobaoFenxiaoProductSkusGetAPIRequest对象
func NewTaobaoFenxiaoProductSkusGetRequest() *TaobaoFenxiaoProductSkusGetAPIRequest {
	return &TaobaoFenxiaoProductSkusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.skus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkusGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetProductId() int64 {
	return r._productId
}
