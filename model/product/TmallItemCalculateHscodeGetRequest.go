package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
算法获取hscode APIRequest
tmall.item.calculate.hscode.get

算法获取hscode
*/
type TmallItemCalculateHscodeGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTmallItemCalculateHscodeGetRequest() *TmallItemCalculateHscodeGetRequest{
    return &TmallItemCalculateHscodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemCalculateHscodeGetRequest) GetApiMethodName() string {
    return "tmall.item.calculate.hscode.get"
}

func (r TmallItemCalculateHscodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemCalculateHscodeGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemCalculateHscodeGetRequest) GetItemId() int64 {
    return r.itemId
}

