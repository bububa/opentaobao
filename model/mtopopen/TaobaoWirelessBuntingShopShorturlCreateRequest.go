package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过店铺id取得短链 API请求
taobao.wireless.bunting.shop.shorturl.create

通过店铺id取得短链
*/
type TaobaoWirelessBuntingShopShorturlCreateRequest struct {
    model.Params
    // 商店id
    _shopId   string
}

// 初始化TaobaoWirelessBuntingShopShorturlCreateRequest对象
func NewTaobaoWirelessBuntingShopShorturlCreateRequest() *TaobaoWirelessBuntingShopShorturlCreateRequest{
    return &TaobaoWirelessBuntingShopShorturlCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWirelessBuntingShopShorturlCreateRequest) GetApiMethodName() string {
    return "taobao.wireless.bunting.shop.shorturl.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWirelessBuntingShopShorturlCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 商店id
func (r *TaobaoWirelessBuntingShopShorturlCreateRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r TaobaoWirelessBuntingShopShorturlCreateRequest) GetShopId() string {
    return r._shopId
}
