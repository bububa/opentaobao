package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单个订阅关系 APIRequest
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationQueryRequest struct {
    model.Params

    // 商品Id
    itemId   int64 

}

func NewTaobaoBaichuanItemSubscribeRelationQueryRequest() *TaobaoBaichuanItemSubscribeRelationQueryRequest{
    return &TaobaoBaichuanItemSubscribeRelationQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBaichuanItemSubscribeRelationQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe.relation.query"
}

func (r TaobaoBaichuanItemSubscribeRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBaichuanItemSubscribeRelationQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoBaichuanItemSubscribeRelationQueryRequest) GetItemId() int64 {
    return r.itemId
}

