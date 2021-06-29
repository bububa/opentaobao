package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询组合商品的组合关系 API请求
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系
*/
type TaobaoWlbWmsItemCombinationGetRequest struct {
    model.Params
    // 货品Id
    itemid   int64
}

// 初始化TaobaoWlbWmsItemCombinationGetRequest对象
func NewTaobaoWlbWmsItemCombinationGetRequest() *TaobaoWlbWmsItemCombinationGetRequest{
    return &TaobaoWlbWmsItemCombinationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsItemCombinationGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.item.combination.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsItemCombinationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Itemid Setter
// 货品Id
func (r *TaobaoWlbWmsItemCombinationGetRequest) SetItemid(itemid int64) error {
    r.itemid = itemid
    r.Set("itemid", itemid)
    return nil
}

// Itemid Getter
func (r TaobaoWlbWmsItemCombinationGetRequest) GetItemid() int64 {
    return r.itemid
}
