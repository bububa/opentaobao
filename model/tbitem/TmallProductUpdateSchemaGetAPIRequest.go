package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductUpdateSchemaGetAPIRequest 产品更新规则获取接口 API请求
// tmall.product.update.schema.get
//
// 获取用户更新产品的规则
type TmallProductUpdateSchemaGetAPIRequest struct {
	model.Params
	// 产品编号
	_productId int64
}

// NewTmallProductUpdateSchemaGetRequest 初始化TmallProductUpdateSchemaGetAPIRequest对象
func NewTmallProductUpdateSchemaGetRequest() *TmallProductUpdateSchemaGetAPIRequest {
	return &TmallProductUpdateSchemaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallProductUpdateSchemaGetAPIRequest) Reset() {
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductUpdateSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductUpdateSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品编号
func (r *TmallProductUpdateSchemaGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallProductUpdateSchemaGetAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTmallProductUpdateSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallProductUpdateSchemaGetRequest()
	},
}

// GetTmallProductUpdateSchemaGetRequest 从 sync.Pool 获取 TmallProductUpdateSchemaGetAPIRequest
func GetTmallProductUpdateSchemaGetAPIRequest() *TmallProductUpdateSchemaGetAPIRequest {
	return poolTmallProductUpdateSchemaGetAPIRequest.Get().(*TmallProductUpdateSchemaGetAPIRequest)
}

// ReleaseTmallProductUpdateSchemaGetAPIRequest 将 TmallProductUpdateSchemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallProductUpdateSchemaGetAPIRequest(v *TmallProductUpdateSchemaGetAPIRequest) {
	v.Reset()
	poolTmallProductUpdateSchemaGetAPIRequest.Put(v)
}
