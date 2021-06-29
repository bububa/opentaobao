package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
打标结果查询-商品维度 API请求
taobao.qimen.items.tag.query

调用该接口，查询某个/某批商品上的标
*/
type TaobaoQimenItemsTagQueryRequest struct {
    model.Params
    // 线上淘宝商品ID，long，必填
    _itemIds   []int64
}

// 初始化TaobaoQimenItemsTagQueryRequest对象
func NewTaobaoQimenItemsTagQueryRequest() *TaobaoQimenItemsTagQueryRequest{
    return &TaobaoQimenItemsTagQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsTagQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.items.tag.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsTagQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 线上淘宝商品ID，long，必填
func (r *TaobaoQimenItemsTagQueryRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoQimenItemsTagQueryRequest) GetItemIds() []int64 {
    return r._itemIds
}
