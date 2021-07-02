package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessBuntingShopShorturlCreateAPIRequest 通过店铺id取得短链 API请求
// taobao.wireless.bunting.shop.shorturl.create
//
// 通过店铺id取得短链
type TaobaoWirelessBuntingShopShorturlCreateAPIRequest struct {
	model.Params
	// 商店id
	_shopId string
}

// NewTaobaoWirelessBuntingShopShorturlCreateRequest 初始化TaobaoWirelessBuntingShopShorturlCreateAPIRequest对象
func NewTaobaoWirelessBuntingShopShorturlCreateRequest() *TaobaoWirelessBuntingShopShorturlCreateAPIRequest {
	return &TaobaoWirelessBuntingShopShorturlCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWirelessBuntingShopShorturlCreateAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.bunting.shop.shorturl.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWirelessBuntingShopShorturlCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetShopId is ShopId Setter
// 商店id
func (r *TaobaoWirelessBuntingShopShorturlCreateAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoWirelessBuntingShopShorturlCreateAPIRequest) GetShopId() string {
	return r._shopId
}
