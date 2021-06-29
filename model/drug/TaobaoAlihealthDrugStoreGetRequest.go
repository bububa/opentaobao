package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据店铺id获取店铺详情 API请求
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情
*/
type TaobaoAlihealthDrugStoreGetRequest struct {
    model.Params
    // 店铺ID
    shopId   int64
}

// 初始化TaobaoAlihealthDrugStoreGetRequest对象
func NewTaobaoAlihealthDrugStoreGetRequest() *TaobaoAlihealthDrugStoreGetRequest{
    return &TaobaoAlihealthDrugStoreGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlihealthDrugStoreGetRequest) GetApiMethodName() string {
    return "taobao.alihealth.drug.store.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlihealthDrugStoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 店铺ID
func (r *TaobaoAlihealthDrugStoreGetRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TaobaoAlihealthDrugStoreGetRequest) GetShopId() int64 {
    return r.shopId
}
