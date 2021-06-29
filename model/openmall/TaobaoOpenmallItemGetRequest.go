package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品详情物料 API请求
taobao.openmall.item.get

获取联盟开放的openmall商品
*/
type TaobaoOpenmallItemGetRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
}

// 初始化TaobaoOpenmallItemGetRequest对象
func NewTaobaoOpenmallItemGetRequest() *TaobaoOpenmallItemGetRequest{
    return &TaobaoOpenmallItemGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallItemGetRequest) GetApiMethodName() string {
    return "taobao.openmall.item.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpenmallItemGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpenmallItemGetRequest) GetItemId() int64 {
    return r._itemId
}
