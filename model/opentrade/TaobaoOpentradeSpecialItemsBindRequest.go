package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单场景商品绑定 APIRequest
taobao.opentrade.special.items.bind

专属下单场景商品绑定
*/
type TaobaoOpentradeSpecialItemsBindRequest struct {
    model.Params

    // 绑定专属下单场景的C端小程序ID
    miniappId   int64 

    // 本次待绑定的商品ID列表
    itemIds   []int64 

}

func NewTaobaoOpentradeSpecialItemsBindRequest() *TaobaoOpentradeSpecialItemsBindRequest{
    return &TaobaoOpentradeSpecialItemsBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeSpecialItemsBindRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.items.bind"
}

func (r TaobaoOpentradeSpecialItemsBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeSpecialItemsBindRequest) SetMiniappId(miniappId int64) error {
    r.miniappId = miniappId
    r.Set("miniapp_id", miniappId)
    return nil
}

func (r TaobaoOpentradeSpecialItemsBindRequest) GetMiniappId() int64 {
    return r.miniappId
}

func (r *TaobaoOpentradeSpecialItemsBindRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoOpentradeSpecialItemsBindRequest) GetItemIds() []int64 {
    return r.itemIds
}

