package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductimagedeleteAPIRequest 产品图片删除 API请求
// taobao.fenxiao.product.image.delete
//
// 产品图片删除，只删除图片信息，不真正删除图片
type TaobaofenxiaoproductimagedeleteAPIRequest struct {
	model.Params
	// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
	_properties string
	// 产品ID
	_productId int64
	// 图片位置
	_position int64
}

// NewTaobaofenxiaoproductimagedeleteRequest 初始化TaobaofenxiaoproductimagedeleteAPIRequest对象
func NewTaobaofenxiaoproductimagedeleteRequest() *TaobaofenxiaoproductimagedeleteAPIRequest {
	return &TaobaofenxiaoproductimagedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductimagedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.image.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductimagedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductimagedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProperties is Properties Setter
// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
func (r *TaobaofenxiaoproductimagedeleteAPIRequest) SetProperties(_properties string) error {
	r._properties = _properties
	r.Set("properties", _properties)
	return nil
}

// GetProperties Properties Getter
func (r TaobaofenxiaoproductimagedeleteAPIRequest) GetProperties() string {
	return r._properties
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaofenxiaoproductimagedeleteAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproductimagedeleteAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetPosition is Position Setter
// 图片位置
func (r *TaobaofenxiaoproductimagedeleteAPIRequest) SetPosition(_position int64) error {
	r._position = _position
	r.Set("position", _position)
	return nil
}

// GetPosition Position Getter
func (r TaobaofenxiaoproductimagedeleteAPIRequest) GetPosition() int64 {
	return r._position
}
