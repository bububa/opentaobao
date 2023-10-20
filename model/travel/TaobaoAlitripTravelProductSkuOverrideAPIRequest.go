package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelproductskuoverrideAPIRequest （供销）产品级别日历价格库存修改，全量覆盖 API请求
// taobao.alitrip.travel.product.sku.override
//
// （供销）产品级别日历价格库存修改，全量覆盖
type TaobaoalitriptravelproductskuoverrideAPIRequest struct {
	model.Params
	// 商品日历价格库存套餐
	_skus []ItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoalitriptravelproductskuoverrideRequest 初始化TaobaoalitriptravelproductskuoverrideAPIRequest对象
func NewTaobaoalitriptravelproductskuoverrideRequest() *TaobaoalitriptravelproductskuoverrideAPIRequest {
	return &TaobaoalitriptravelproductskuoverrideAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelproductskuoverrideAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.product.sku.override"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelproductskuoverrideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelproductskuoverrideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkus is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoalitriptravelproductskuoverrideAPIRequest) SetSkus(_skus []ItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// GetSkus Skus Getter
func (r TaobaoalitriptravelproductskuoverrideAPIRequest) GetSkus() []ItemSkuInfo {
	return r._skus
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelproductskuoverrideAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoalitriptravelproductskuoverrideAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelproductskuoverrideAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoalitriptravelproductskuoverrideAPIRequest) GetItemId() int64 {
	return r._itemId
}
