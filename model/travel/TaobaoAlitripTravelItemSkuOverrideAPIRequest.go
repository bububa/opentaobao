package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemskuoverrideAPIRequest 【API3.0】商品级别日历价格库存修改，全量覆盖 API请求
// taobao.alitrip.travel.item.sku.override
//
// 旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
type TaobaoalitriptravelitemskuoverrideAPIRequest struct {
	model.Params
	// 商品日历价格库存套餐
	_skus []PontusTravelItemSkuInfo
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
}

// NewTaobaoalitriptravelitemskuoverrideRequest 初始化TaobaoalitriptravelitemskuoverrideAPIRequest对象
func NewTaobaoalitriptravelitemskuoverrideRequest() *TaobaoalitriptravelitemskuoverrideAPIRequest {
	return &TaobaoalitriptravelitemskuoverrideAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelitemskuoverrideAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.sku.override"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelitemskuoverrideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelitemskuoverrideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkus is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoalitriptravelitemskuoverrideAPIRequest) SetSkus(_skus []PontusTravelItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// GetSkus Skus Getter
func (r TaobaoalitriptravelitemskuoverrideAPIRequest) GetSkus() []PontusTravelItemSkuInfo {
	return r._skus
}

// SetOutProductId is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelitemskuoverrideAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r TaobaoalitriptravelitemskuoverrideAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoalitriptravelitemskuoverrideAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoalitriptravelitemskuoverrideAPIRequest) GetItemId() int64 {
	return r._itemId
}
