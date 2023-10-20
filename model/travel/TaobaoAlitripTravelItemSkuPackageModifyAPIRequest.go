package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemskupackagemodifyAPIRequest 【API3.0】套餐级别日历价格库存增删操作 API请求
// taobao.alitrip.travel.item.sku.package.modify
//
// 【API3.0】套餐级别日历价格库存增删操作
type TaobaoalitriptravelitemskupackagemodifyAPIRequest struct {
	model.Params
	// 商品日历价格库存套餐
	_skus []ItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoalitriptravelitemskupackagemodifyRequest 初始化TaobaoalitriptravelitemskupackagemodifyAPIRequest对象
func NewTaobaoalitriptravelitemskupackagemodifyRequest() *TaobaoalitriptravelitemskupackagemodifyAPIRequest {
	return &TaobaoalitriptravelitemskupackagemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelitemskupackagemodifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.sku.package.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelitemskupackagemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelitemskupackagemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkus is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoalitriptravelitemskupackagemodifyAPIRequest) SetSkus(_skus []ItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// GetSkus Skus Getter
func (r TaobaoalitriptravelitemskupackagemodifyAPIRequest) GetSkus() []ItemSkuInfo {
	return r._skus
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelitemskupackagemodifyAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoalitriptravelitemskupackagemodifyAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelitemskupackagemodifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoalitriptravelitemskupackagemodifyAPIRequest) GetItemId() int64 {
	return r._itemId
}
