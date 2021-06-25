package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品上架 APIRequest
alibaba.item.operate.upshelf

商品上架
*/
type AlibabaItemOperateUpshelfRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

    // 商品库存
    quantity   int64 

}

func NewAlibabaItemOperateUpshelfRequest() *AlibabaItemOperateUpshelfRequest{
    return &AlibabaItemOperateUpshelfRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemOperateUpshelfRequest) GetApiMethodName() string {
    return "alibaba.item.operate.upshelf"
}

func (r AlibabaItemOperateUpshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemOperateUpshelfRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaItemOperateUpshelfRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaItemOperateUpshelfRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

func (r AlibabaItemOperateUpshelfRequest) GetQuantity() int64 {
    return r.quantity
}

