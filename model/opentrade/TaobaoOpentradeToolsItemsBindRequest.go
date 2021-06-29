package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易开放商品绑定 APIRequest
taobao.opentrade.tools.items.bind

交易开放商品绑定
*/
type TaobaoOpentradeToolsItemsBindRequest struct {
    model.Params

    // 绑定交易开放场景的C端小程序ID
    miniappId   int64 

    // 待绑定商品id
    itemIds   []int64 

}

func NewTaobaoOpentradeToolsItemsBindRequest() *TaobaoOpentradeToolsItemsBindRequest{
    return &TaobaoOpentradeToolsItemsBindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeToolsItemsBindRequest) GetApiMethodName() string {
    return "taobao.opentrade.tools.items.bind"
}

func (r TaobaoOpentradeToolsItemsBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeToolsItemsBindRequest) SetMiniappId(miniappId int64) error {
    r.miniappId = miniappId
    r.Set("miniapp_id", miniappId)
    return nil
}

func (r TaobaoOpentradeToolsItemsBindRequest) GetMiniappId() int64 {
    return r.miniappId
}

func (r *TaobaoOpentradeToolsItemsBindRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoOpentradeToolsItemsBindRequest) GetItemIds() []int64 {
    return r.itemIds
}

