package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫品牌新品同步API API请求
tmall.brand.item.upload

支撑天猫品牌将各渠道新品信息同步至平台
*/
type TmallBrandItemUploadAPIRequest struct {
    model.Params
    // 需要同步的商品列表
    _itemList   []TmallBrandChannelNewItem
}

// 初始化TmallBrandItemUploadAPIRequest对象
func NewTmallBrandItemUploadRequest() *TmallBrandItemUploadAPIRequest{
    return &TmallBrandItemUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallBrandItemUploadAPIRequest) GetApiMethodName() string {
    return "tmall.brand.item.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallBrandItemUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemList Setter
// 需要同步的商品列表
func (r *TmallBrandItemUploadAPIRequest) SetItemList(_itemList []TmallBrandChannelNewItem) error {
    r._itemList = _itemList
    r.Set("item_list", _itemList)
    return nil
}

// ItemList Getter
func (r TmallBrandItemUploadAPIRequest) GetItemList() []TmallBrandChannelNewItem {
    return r._itemList
}
