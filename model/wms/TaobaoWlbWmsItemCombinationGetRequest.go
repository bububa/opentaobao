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
type TaobaoWlbWmsItemCombinationGetAPIRequest struct {
    model.Params
    // 货品Id
    _itemid   int64
}

// 初始化TaobaoWlbWmsItemCombinationGetAPIRequest对象
func NewTaobaoWlbWmsItemCombinationGetRequest() *TaobaoWlbWmsItemCombinationGetAPIRequest{
    return &TaobaoWlbWmsItemCombinationGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.item.combination.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Itemid Setter
// 货品Id
func (r *TaobaoWlbWmsItemCombinationGetAPIRequest) SetItemid(_itemid int64) error {
    r._itemid = _itemid
    r.Set("itemid", _itemid)
    return nil
}

// Itemid Getter
func (r TaobaoWlbWmsItemCombinationGetAPIRequest) GetItemid() int64 {
    return r._itemid
}
