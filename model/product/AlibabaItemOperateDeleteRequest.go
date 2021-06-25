package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品删除 APIRequest
alibaba.item.operate.delete

商品删除
*/
type AlibabaItemOperateDeleteRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

}

func NewAlibabaItemOperateDeleteRequest() *AlibabaItemOperateDeleteRequest{
    return &AlibabaItemOperateDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItemOperateDeleteRequest) GetApiMethodName() string {
    return "alibaba.item.operate.delete"
}

func (r AlibabaItemOperateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItemOperateDeleteRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaItemOperateDeleteRequest) GetItemId() int64 {
    return r.itemId
}

