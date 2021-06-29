package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单个订阅关系 API请求
taobao.baichuan.item.subscribe.relation.query

查询单个订阅关系
*/
type TaobaoBaichuanItemSubscribeRelationQueryRequest struct {
    model.Params
    // 商品Id
    _itemId   int64
}

// 初始化TaobaoBaichuanItemSubscribeRelationQueryRequest对象
func NewTaobaoBaichuanItemSubscribeRelationQueryRequest() *TaobaoBaichuanItemSubscribeRelationQueryRequest{
    return &TaobaoBaichuanItemSubscribeRelationQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanItemSubscribeRelationQueryRequest) GetApiMethodName() string {
    return "taobao.baichuan.item.subscribe.relation.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanItemSubscribeRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品Id
func (r *TaobaoBaichuanItemSubscribeRelationQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBaichuanItemSubscribeRelationQueryRequest) GetItemId() int64 {
    return r._itemId
}
