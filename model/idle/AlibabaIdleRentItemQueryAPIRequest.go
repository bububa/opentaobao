package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询租赁商品信息 API请求
alibaba.idle.rent.item.query

查询租赁商品信息
*/
type AlibabaIdleRentItemQueryAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化AlibabaIdleRentItemQueryAPIRequest对象
func NewAlibabaIdleRentItemQueryRequest() *AlibabaIdleRentItemQueryAPIRequest{
    return &AlibabaIdleRentItemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *AlibabaIdleRentItemQueryAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaIdleRentItemQueryAPIRequest) GetItemId() int64 {
    return r._itemId
}
