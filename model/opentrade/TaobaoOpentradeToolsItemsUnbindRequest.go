package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易开放商品解绑 APIRequest
taobao.opentrade.tools.items.unbind

交易开放商品解绑
*/
type TaobaoOpentradeToolsItemsUnbindRequest struct {
    model.Params

    // 绑定交易开放场景的C端小程序ID
    miniappId   int64 

    // 商品id
    itemIds   []int64 

}

func NewTaobaoOpentradeToolsItemsUnbindRequest() *TaobaoOpentradeToolsItemsUnbindRequest{
    return &TaobaoOpentradeToolsItemsUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeToolsItemsUnbindRequest) GetApiMethodName() string {
    return "taobao.opentrade.tools.items.unbind"
}

func (r TaobaoOpentradeToolsItemsUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeToolsItemsUnbindRequest) SetMiniappId(miniappId int64) error {
    r.miniappId = miniappId
    r.Set("miniapp_id", miniappId)
    return nil
}

func (r TaobaoOpentradeToolsItemsUnbindRequest) GetMiniappId() int64 {
    return r.miniappId
}

func (r *TaobaoOpentradeToolsItemsUnbindRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoOpentradeToolsItemsUnbindRequest) GetItemIds() []int64 {
    return r.itemIds
}

