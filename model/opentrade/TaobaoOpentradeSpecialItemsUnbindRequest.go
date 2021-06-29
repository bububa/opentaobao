package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单场景商品解绑 APIRequest
taobao.opentrade.special.items.unbind

专属下单场景商品解绑
*/
type TaobaoOpentradeSpecialItemsUnbindRequest struct {
    model.Params

    // 绑定专属下单场景的C端小程序ID
    miniappId   int64 

    // 本次待解绑的商品ID列表
    itemIds   []int64 

}

func NewTaobaoOpentradeSpecialItemsUnbindRequest() *TaobaoOpentradeSpecialItemsUnbindRequest{
    return &TaobaoOpentradeSpecialItemsUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeSpecialItemsUnbindRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.items.unbind"
}

func (r TaobaoOpentradeSpecialItemsUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeSpecialItemsUnbindRequest) SetMiniappId(miniappId int64) error {
    r.miniappId = miniappId
    r.Set("miniapp_id", miniappId)
    return nil
}

func (r TaobaoOpentradeSpecialItemsUnbindRequest) GetMiniappId() int64 {
    return r.miniappId
}

func (r *TaobaoOpentradeSpecialItemsUnbindRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoOpentradeSpecialItemsUnbindRequest) GetItemIds() []int64 {
    return r.itemIds
}

