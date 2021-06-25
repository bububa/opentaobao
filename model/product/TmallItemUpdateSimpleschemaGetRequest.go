package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网同购编辑商品的get接口 APIRequest
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口
*/
type TmallItemUpdateSimpleschemaGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTmallItemUpdateSimpleschemaGetRequest() *TmallItemUpdateSimpleschemaGetRequest{
    return &TmallItemUpdateSimpleschemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemUpdateSimpleschemaGetRequest) GetApiMethodName() string {
    return "tmall.item.update.simpleschema.get"
}

func (r TmallItemUpdateSimpleschemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemUpdateSimpleschemaGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemUpdateSimpleschemaGetRequest) GetItemId() int64 {
    return r.itemId
}

