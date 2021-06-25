package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品获取接口 APIRequest
tmall.item.combine.get

查询组合商品的SKU信息
*/
type TmallItemCombineGetRequest struct {
    model.Params

    // 组合商品ID
    itemId   int64 

}

func NewTmallItemCombineGetRequest() *TmallItemCombineGetRequest{
    return &TmallItemCombineGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemCombineGetRequest) GetApiMethodName() string {
    return "tmall.item.combine.get"
}

func (r TmallItemCombineGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemCombineGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemCombineGetRequest) GetItemId() int64 {
    return r.itemId
}

