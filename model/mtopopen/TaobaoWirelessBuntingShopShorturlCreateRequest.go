package mtopopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过店铺id取得短链 APIRequest
taobao.wireless.bunting.shop.shorturl.create

通过店铺id取得短链
*/
type TaobaoWirelessBuntingShopShorturlCreateRequest struct {
    model.Params

    // 商店id
    shopId   string 

}

func NewTaobaoWirelessBuntingShopShorturlCreateRequest() *TaobaoWirelessBuntingShopShorturlCreateRequest{
    return &TaobaoWirelessBuntingShopShorturlCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWirelessBuntingShopShorturlCreateRequest) GetApiMethodName() string {
    return "taobao.wireless.bunting.shop.shorturl.create"
}

func (r TaobaoWirelessBuntingShopShorturlCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWirelessBuntingShopShorturlCreateRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TaobaoWirelessBuntingShopShorturlCreateRequest) GetShopId() string {
    return r.shopId
}

