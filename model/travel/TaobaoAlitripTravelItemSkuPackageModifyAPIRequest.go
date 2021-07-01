package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemSkuPackageModifyAPIRequest
【API3.0】套餐级别日历价格库存增删操作 API请求
taobao.alitrip.travel.item.sku.package.modify

【API3.0】套餐级别日历价格库存增删操作 */
type TaobaoAlitripTravelItemSkuPackageModifyAPIRequest struct {
	model.Params
	// 商品id。itemId和outProductId至少填写一个
	_itemId int64
	// 商品 外部商家编码。itemId和outProductId至少填写一个
	_outProductId string
	// 商品日历价格库存套餐
	_skus []ItemSkuInfo
}

// NewTaobaoAlitripTravelItemSkuPackageModifyRequest 初始化TaobaoAlitripTravelItemSkuPackageModifyAPIRequest对象
func NewTaobaoAlitripTravelItemSkuPackageModifyRequest() *TaobaoAlitripTravelItemSkuPackageModifyAPIRequest {
	return &TaobaoAlitripTravelItemSkuPackageModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.sku.package.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// Get OutProductId Getter
func (r TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// Set is Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) SetSkus(_skus []ItemSkuInfo) error {
	r._skus = _skus
	r.Set("skus", _skus)
	return nil
}

// Get Skus Getter
func (r TaobaoAlitripTravelItemSkuPackageModifyAPIRequest) GetSkus() []ItemSkuInfo {
	return r._skus
}
