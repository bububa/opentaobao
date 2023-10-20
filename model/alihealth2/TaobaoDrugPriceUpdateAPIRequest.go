package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodrugpriceupdateAPIRequest 商家更新宝贝价格 API请求
// taobao.drug.price.update
//
// 商家更新价格
type TaobaodrugpriceupdateAPIRequest struct {
	model.Params
	// 对应的外部店铺ID
	_outStoreId string
	// 对应的外部商品编码
	_outItemId string
	// 商品价格
	_price float64
}

// NewTaobaodrugpriceupdateRequest 初始化TaobaodrugpriceupdateAPIRequest对象
func NewTaobaodrugpriceupdateRequest() *TaobaodrugpriceupdateAPIRequest {
	return &TaobaodrugpriceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodrugpriceupdateAPIRequest) GetApiMethodName() string {
	return "taobao.drug.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodrugpriceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodrugpriceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutStoreId is OutStoreId Setter
// 对应的外部店铺ID
func (r *TaobaodrugpriceupdateAPIRequest) SetOutStoreId(_outStoreId string) error {
	r._outStoreId = _outStoreId
	r.Set("out_store_id", _outStoreId)
	return nil
}

// GetOutStoreId OutStoreId Getter
func (r TaobaodrugpriceupdateAPIRequest) GetOutStoreId() string {
	return r._outStoreId
}

// SetOutItemId is OutItemId Setter
// 对应的外部商品编码
func (r *TaobaodrugpriceupdateAPIRequest) SetOutItemId(_outItemId string) error {
	r._outItemId = _outItemId
	r.Set("out_item_id", _outItemId)
	return nil
}

// GetOutItemId OutItemId Getter
func (r TaobaodrugpriceupdateAPIRequest) GetOutItemId() string {
	return r._outItemId
}

// SetPrice is Price Setter
// 商品价格
func (r *TaobaodrugpriceupdateAPIRequest) SetPrice(_price float64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaodrugpriceupdateAPIRequest) GetPrice() float64 {
	return r._price
}
