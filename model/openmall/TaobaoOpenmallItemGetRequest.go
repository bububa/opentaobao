package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品详情物料 APIRequest
taobao.openmall.item.get

获取联盟开放的openmall商品
*/
type TaobaoOpenmallItemGetRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

}

func NewTaobaoOpenmallItemGetRequest() *TaobaoOpenmallItemGetRequest{
    return &TaobaoOpenmallItemGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallItemGetRequest) GetApiMethodName() string {
    return "taobao.openmall.item.get"
}

func (r TaobaoOpenmallItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallItemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOpenmallItemGetRequest) GetItemId() int64 {
    return r.itemId
}

