package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 APIRequest
alibaba.item.operate.downshelf

商品下架
*/
type AlibabaItemOperateDownshelfRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

}

func NewAlibabaItemOperateDownshelfRequest() *AlibabaItemOperateDownshelfRequest{
    return &AlibabaItemOperateDownshelfRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemOperateDownshelfRequest) GetApiMethodName() string {
    return "alibaba.item.operate.downshelf"
}

func (r AlibabaItemOperateDownshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemOperateDownshelfRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaItemOperateDownshelfRequest) GetItemId() int64 {
    return r.itemId
}

