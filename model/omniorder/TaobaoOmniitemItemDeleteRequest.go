package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品删除 APIRequest
taobao.omniitem.item.delete

全渠道商品删除，能够对门店商品库商品进行删除动作
*/
type TaobaoOmniitemItemDeleteRequest struct {
    model.Params

    // 条形码
    barCode   string 

    // 商品ID，若填入则以该字段为准，否则以outerId+barcode为准
    itemId   int64 

    // 商品outerId
    outerId   string 

}

func NewTaobaoOmniitemItemDeleteRequest() *TaobaoOmniitemItemDeleteRequest{
    return &TaobaoOmniitemItemDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemItemDeleteRequest) GetApiMethodName() string {
    return "taobao.omniitem.item.delete"
}

func (r TaobaoOmniitemItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemItemDeleteRequest) SetBarCode(barCode string) error {
    r.barCode = barCode
    r.Set("bar_code", barCode)
    return nil
}

func (r TaobaoOmniitemItemDeleteRequest) GetBarCode() string {
    return r.barCode
}

func (r *TaobaoOmniitemItemDeleteRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOmniitemItemDeleteRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOmniitemItemDeleteRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoOmniitemItemDeleteRequest) GetOuterId() string {
    return r.outerId
}

