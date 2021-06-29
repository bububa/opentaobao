package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝商品ID查询商品映射关系 API请求
taobao.wlb.item.map.get

根据物流宝商品ID查询商品映射关系
*/
type TaobaoWlbItemMapGetRequest struct {
    model.Params
    // 要查询映射关系的物流宝商品id
    _itemId   int64
}

// 初始化TaobaoWlbItemMapGetRequest对象
func NewTaobaoWlbItemMapGetRequest() *TaobaoWlbItemMapGetRequest{
    return &TaobaoWlbItemMapGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbItemMapGetRequest) GetApiMethodName() string {
    return "taobao.wlb.item.map.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbItemMapGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 要查询映射关系的物流宝商品id
func (r *TaobaoWlbItemMapGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoWlbItemMapGetRequest) GetItemId() int64 {
    return r._itemId
}
