package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
vip商家编辑商品的规则获取接口 APIRequest
tmall.item.vip.update.schema.get

获取vip商家编辑商品的规则
*/
type TmallItemVipUpdateSchemaGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTmallItemVipUpdateSchemaGetRequest() *TmallItemVipUpdateSchemaGetRequest{
    return &TmallItemVipUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemVipUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.vip.update.schema.get"
}

func (r TmallItemVipUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemVipUpdateSchemaGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemVipUpdateSchemaGetRequest) GetItemId() int64 {
    return r.itemId
}

