package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductskudeleteAPIRequest 产品SKU删除接口 API请求
// taobao.fenxiao.product.sku.delete
//
// 根据sku properties删除sku数据
type TaobaofenxiaoproductskudeleteAPIRequest struct {
	model.Params
	// sku属性
	_properties string
	// 产品id
	_productId int64
}

// NewTaobaofenxiaoproductskudeleteRequest 初始化TaobaofenxiaoproductskudeleteAPIRequest对象
func NewTaobaofenxiaoproductskudeleteRequest() *TaobaofenxiaoproductskudeleteAPIRequest {
	return &TaobaofenxiaoproductskudeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductskudeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductskudeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductskudeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// sku属性
func (r *TaobaofenxiaoproductskudeleteAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaofenxiaoproductskudeleteAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaofenxiaoproductskudeleteAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductskudeleteAPIRequest) GetProductId() int64 {
	return r._productId
}
