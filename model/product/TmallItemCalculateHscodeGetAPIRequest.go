package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
算法获取hscode API请求
tmall.item.calculate.hscode.get

算法获取hscode
*/
type TmallItemCalculateHscodeGetAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TmallItemCalculateHscodeGetAPIRequest对象
func NewTmallItemCalculateHscodeGetRequest() *TmallItemCalculateHscodeGetAPIRequest{
    return &TmallItemCalculateHscodeGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemCalculateHscodeGetAPIRequest) GetApiMethodName() string {
    return "tmall.item.calculate.hscode.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemCalculateHscodeGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallItemCalculateHscodeGetAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemCalculateHscodeGetAPIRequest) GetItemId() int64 {
    return r._itemId
}
