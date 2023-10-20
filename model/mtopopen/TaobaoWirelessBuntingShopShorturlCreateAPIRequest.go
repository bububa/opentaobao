package mtopopen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWirelessBuntingShopShorturlCreateAPIRequest) Reset() {
	r._shopId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWirelessBuntingShopShorturlCreateAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.bunting.shop.shorturl.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWirelessBuntingShopShorturlCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWirelessBuntingShopShorturlCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoWirelessBuntingShopShorturlCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWirelessBuntingShopShorturlCreateRequest()
	},
}

// GetTaobaoWirelessBuntingShopShorturlCreateRequest 从 sync.Pool 获取 TaobaoWirelessBuntingShopShorturlCreateAPIRequest
func GetTaobaoWirelessBuntingShopShorturlCreateAPIRequest() *TaobaoWirelessBuntingShopShorturlCreateAPIRequest {
	return poolTaobaoWirelessBuntingShopShorturlCreateAPIRequest.Get().(*TaobaoWirelessBuntingShopShorturlCreateAPIRequest)
}

// ReleaseTaobaoWirelessBuntingShopShorturlCreateAPIRequest 将 TaobaoWirelessBuntingShopShorturlCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoWirelessBuntingShopShorturlCreateAPIRequest(v *TaobaoWirelessBuntingShopShorturlCreateAPIRequest) {
	v.Reset()
	poolTaobaoWirelessBuntingShopShorturlCreateAPIRequest.Put(v)
}
