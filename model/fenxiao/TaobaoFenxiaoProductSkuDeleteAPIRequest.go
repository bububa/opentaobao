package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductSkuDeleteAPIRequest 产品SKU删除接口 API请求
// taobao.fenxiao.product.sku.delete
//
// 根据sku properties删除sku数据
type TaobaoFenxiaoProductSkuDeleteAPIRequest struct {
	model.Params
	// sku属性
	_properties string
	// 产品id
	_productId int64
}

// NewTaobaoFenxiaoProductSkuDeleteRequest 初始化TaobaoFenxiaoProductSkuDeleteAPIRequest对象
func NewTaobaoFenxiaoProductSkuDeleteRequest() *TaobaoFenxiaoProductSkuDeleteAPIRequest {
	return &TaobaoFenxiaoProductSkuDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductSkuDeleteAPIRequest) Reset() {
	r._properties = ""
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuDeleteAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaoFenxiaoProductSkuDeleteAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductSkuDeleteAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTaobaoFenxiaoProductSkuDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductSkuDeleteRequest()
	},
}

// GetTaobaoFenxiaoProductSkuDeleteRequest 从 sync.Pool 获取 TaobaoFenxiaoProductSkuDeleteAPIRequest
func GetTaobaoFenxiaoProductSkuDeleteAPIRequest() *TaobaoFenxiaoProductSkuDeleteAPIRequest {
	return poolTaobaoFenxiaoProductSkuDeleteAPIRequest.Get().(*TaobaoFenxiaoProductSkuDeleteAPIRequest)
}

// ReleaseTaobaoFenxiaoProductSkuDeleteAPIRequest 将 TaobaoFenxiaoProductSkuDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductSkuDeleteAPIRequest(v *TaobaoFenxiaoProductSkuDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductSkuDeleteAPIRequest.Put(v)
}
