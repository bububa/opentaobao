package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品删除 API请求
alibaba.item.operate.delete

商品删除
*/
type AlibabaItemOperateDeleteRequest struct {
    model.Params
    // 商品ID
    itemId   int64
}

// 初始化AlibabaItemOperateDeleteRequest对象
func NewAlibabaItemOperateDeleteRequest() *AlibabaItemOperateDeleteRequest{
    return &AlibabaItemOperateDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItemOperateDeleteRequest) GetApiMethodName() string {
    return "alibaba.item.operate.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItemOperateDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *AlibabaItemOperateDeleteRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaItemOperateDeleteRequest) GetItemId() int64 {
    return r.itemId
}
