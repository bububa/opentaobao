package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
vip商家编辑商品的规则获取接口 API请求
tmall.item.vip.update.schema.get

获取vip商家编辑商品的规则
*/
type TmallItemVipUpdateSchemaGetRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TmallItemVipUpdateSchemaGetRequest对象
func NewTmallItemVipUpdateSchemaGetRequest() *TmallItemVipUpdateSchemaGetRequest{
    return &TmallItemVipUpdateSchemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemVipUpdateSchemaGetRequest) GetApiMethodName() string {
    return "tmall.item.vip.update.schema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemVipUpdateSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallItemVipUpdateSchemaGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemVipUpdateSchemaGetRequest) GetItemId() int64 {
    return r._itemId
}
