package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 API请求
alibaba.item.operate.downshelf

商品下架
*/
type AlibabaItemOperateDownshelfRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
}

// 初始化AlibabaItemOperateDownshelfRequest对象
func NewAlibabaItemOperateDownshelfRequest() *AlibabaItemOperateDownshelfRequest{
    return &AlibabaItemOperateDownshelfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemOperateDownshelfRequest) GetApiMethodName() string {
    return "alibaba.item.operate.downshelf"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemOperateDownshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemOperateDownshelfRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemOperateDownshelfRequest) GetItemId() int64 {
    return r._itemId
}
