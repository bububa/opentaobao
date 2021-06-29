package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品获取接口 API请求
tmall.item.combine.get

查询组合商品的SKU信息
*/
type TmallItemCombineGetRequest struct {
    model.Params
    // 组合商品ID
    _itemId   int64
}

// 初始化TmallItemCombineGetRequest对象
func NewTmallItemCombineGetRequest() *TmallItemCombineGetRequest{
    return &TmallItemCombineGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemCombineGetRequest) GetApiMethodName() string {
    return "tmall.item.combine.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemCombineGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 组合商品ID
func (r *TmallItemCombineGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemCombineGetRequest) GetItemId() int64 {
    return r._itemId
}
