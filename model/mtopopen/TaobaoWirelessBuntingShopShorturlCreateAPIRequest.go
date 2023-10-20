package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowirelessbuntingshopshorturlcreateAPIRequest 通过店铺id取得短链 API请求
// taobao.wireless.bunting.shop.shorturl.create
//
// 通过店铺id取得短链
type TaobaowirelessbuntingshopshorturlcreateAPIRequest struct {
	model.Params
	// 商店id
	_shopId string
}

// NewTaobaowirelessbuntingshopshorturlcreateRequest 初始化TaobaowirelessbuntingshopshorturlcreateAPIRequest对象
func NewTaobaowirelessbuntingshopshorturlcreateRequest() *TaobaowirelessbuntingshopshorturlcreateAPIRequest {
	return &TaobaowirelessbuntingshopshorturlcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowirelessbuntingshopshorturlcreateAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.bunting.shop.shorturl.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowirelessbuntingshopshorturlcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowirelessbuntingshopshorturlcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 商店id
func (r *TaobaowirelessbuntingshopshorturlcreateAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaowirelessbuntingshopshorturlcreateAPIRequest) GetShopId() string {
	return r._shopId
}
