package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品上架 API请求
alibaba.item.operate.upshelf

商品上架
*/
type AlibabaItemOperateUpshelfRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 商品库存
    _quantity   int64
}

// 初始化AlibabaItemOperateUpshelfRequest对象
func NewAlibabaItemOperateUpshelfRequest() *AlibabaItemOperateUpshelfRequest{
    return &AlibabaItemOperateUpshelfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemOperateUpshelfRequest) GetApiMethodName() string {
    return "alibaba.item.operate.upshelf"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemOperateUpshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemOperateUpshelfRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemOperateUpshelfRequest) GetItemId() int64 {
    return r._itemId
}
// Quantity Setter
// 商品库存
func (r *AlibabaItemOperateUpshelfRequest) SetQuantity(_quantity int64) error {
    r._quantity = _quantity
    r.Set("quantity", _quantity)
    return nil
}

// Quantity Getter
func (r AlibabaItemOperateUpshelfRequest) GetQuantity() int64 {
    return r._quantity
}
