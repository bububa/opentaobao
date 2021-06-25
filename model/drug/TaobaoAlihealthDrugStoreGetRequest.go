package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据店铺id获取店铺详情 APIRequest
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情
*/
type TaobaoAlihealthDrugStoreGetRequest struct {
    model.Params

    // 店铺ID
    shopId   int64 

}

func NewTaobaoAlihealthDrugStoreGetRequest() *TaobaoAlihealthDrugStoreGetRequest{
    return &TaobaoAlihealthDrugStoreGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlihealthDrugStoreGetRequest) GetApiMethodName() string {
    return "taobao.alihealth.drug.store.get"
}

func (r TaobaoAlihealthDrugStoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlihealthDrugStoreGetRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TaobaoAlihealthDrugStoreGetRequest) GetShopId() int64 {
    return r.shopId
}

