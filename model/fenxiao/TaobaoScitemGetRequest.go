package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询商品 APIRequest
taobao.scitem.get

根据id查询商品
*/
type TaobaoScitemGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTaobaoScitemGetRequest() *TaobaoScitemGetRequest{
    return &TaobaoScitemGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoScitemGetRequest) GetApiMethodName() string {
    return "taobao.scitem.get"
}

func (r TaobaoScitemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoScitemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoScitemGetRequest) GetItemId() int64 {
    return r.itemId
}

