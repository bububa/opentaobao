package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductSkusGetAPIRequest) Reset() {
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.skus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkusGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductSkusGetAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTaobaoFenxiaoProductSkusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductSkusGetRequest()
	},
}

// GetTaobaoFenxiaoProductSkusGetRequest 从 sync.Pool 获取 TaobaoFenxiaoProductSkusGetAPIRequest
func GetTaobaoFenxiaoProductSkusGetAPIRequest() *TaobaoFenxiaoProductSkusGetAPIRequest {
	return poolTaobaoFenxiaoProductSkusGetAPIRequest.Get().(*TaobaoFenxiaoProductSkusGetAPIRequest)
}

// ReleaseTaobaoFenxiaoProductSkusGetAPIRequest 将 TaobaoFenxiaoProductSkusGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductSkusGetAPIRequest(v *TaobaoFenxiaoProductSkusGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductSkusGetAPIRequest.Put(v)
}
