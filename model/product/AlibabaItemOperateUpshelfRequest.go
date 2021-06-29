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
    itemId   int64
    // 商品库存
    quantity   int64
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
func (r *AlibabaItemOperateUpshelfRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemOperateUpshelfRequest) GetItemId() int64 {
    return r.itemId
}
// Quantity Setter
// 商品库存
func (r *AlibabaItemOperateUpshelfRequest) SetQuantity(quantity int64) error {
    r.quantity = quantity
    r.Set("quantity", quantity)
    return nil
}

// Quantity Getter
func (r AlibabaItemOperateUpshelfRequest) GetQuantity() int64 {
    return r.quantity
}
