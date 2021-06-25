package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过商品ID获取标签列表 APIRequest
tmall.traderate.itemtags.get

通过商品ID获取标签详细信息
*/
type TmallTraderateItemtagsGetRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

}

func NewTmallTraderateItemtagsGetRequest() *TmallTraderateItemtagsGetRequest{
    return &TmallTraderateItemtagsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraderateItemtagsGetRequest) GetApiMethodName() string {
    return "tmall.traderate.itemtags.get"
}

func (r TmallTraderateItemtagsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraderateItemtagsGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallTraderateItemtagsGetRequest) GetItemId() int64 {
    return r.itemId
}

