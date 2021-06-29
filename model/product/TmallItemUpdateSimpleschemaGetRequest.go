package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网同购编辑商品的get接口 API请求
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口
*/
type TmallItemUpdateSimpleschemaGetRequest struct {
    model.Params
    // 商品id
    itemId   int64
}

// 初始化TmallItemUpdateSimpleschemaGetRequest对象
func NewTmallItemUpdateSimpleschemaGetRequest() *TmallItemUpdateSimpleschemaGetRequest{
    return &TmallItemUpdateSimpleschemaGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemUpdateSimpleschemaGetRequest) GetApiMethodName() string {
    return "tmall.item.update.simpleschema.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemUpdateSimpleschemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallItemUpdateSimpleschemaGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallItemUpdateSimpleschemaGetRequest) GetItemId() int64 {
    return r.itemId
}
