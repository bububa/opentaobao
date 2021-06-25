package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询组合商品的组合关系 APIRequest
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系
*/
type TaobaoWlbWmsItemCombinationGetRequest struct {
    model.Params

    // 货品Id
    itemid   int64 

}

func NewTaobaoWlbWmsItemCombinationGetRequest() *TaobaoWlbWmsItemCombinationGetRequest{
    return &TaobaoWlbWmsItemCombinationGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsItemCombinationGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.item.combination.get"
}

func (r TaobaoWlbWmsItemCombinationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsItemCombinationGetRequest) SetItemid(itemid int64) error {
    r.itemid = itemid
    r.Set("itemid", itemid)
    return nil
}

func (r TaobaoWlbWmsItemCombinationGetRequest) GetItemid() int64 {
    return r.itemid
}

