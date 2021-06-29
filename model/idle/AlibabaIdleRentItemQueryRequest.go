package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询租赁商品信息 APIRequest
alibaba.idle.rent.item.query

查询租赁商品信息
*/
type AlibabaIdleRentItemQueryRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewAlibabaIdleRentItemQueryRequest() *AlibabaIdleRentItemQueryRequest{
    return &AlibabaIdleRentItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRentItemQueryRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.query"
}

func (r AlibabaIdleRentItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRentItemQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaIdleRentItemQueryRequest) GetItemId() int64 {
    return r.itemId
}

