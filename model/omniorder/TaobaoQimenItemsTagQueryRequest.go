package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
打标结果查询-商品维度 APIRequest
taobao.qimen.items.tag.query

调用该接口，查询某个/某批商品上的标
*/
type TaobaoQimenItemsTagQueryRequest struct {
    model.Params

    // 线上淘宝商品ID，long，必填
    itemIds   []int64 

}

func NewTaobaoQimenItemsTagQueryRequest() *TaobaoQimenItemsTagQueryRequest{
    return &TaobaoQimenItemsTagQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemsTagQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.items.tag.query"
}

func (r TaobaoQimenItemsTagQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemsTagQueryRequest) SetItemIds(itemIds []int64) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoQimenItemsTagQueryRequest) GetItemIds() []int64 {
    return r.itemIds
}

