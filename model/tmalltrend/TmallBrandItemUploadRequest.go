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
type TmallBrandItemUploadRequest struct {
    model.Params
    // 需要同步的商品列表
    _itemList   []TmallBrandChannelNewItem
}

// 初始化TmallBrandItemUploadRequest对象
func NewTmallBrandItemUploadRequest() *TmallBrandItemUploadRequest{
    return &TmallBrandItemUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallBrandItemUploadRequest) GetApiMethodName() string {
    return "tmall.brand.item.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallBrandItemUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemList Setter
// 需要同步的商品列表
func (r *TmallBrandItemUploadRequest) SetItemList(_itemList []TmallBrandChannelNewItem) error {
    r._itemList = _itemList
    r.Set("item_list", _itemList)
    return nil
}

// ItemList Getter
func (r TmallBrandItemUploadRequest) GetItemList() []TmallBrandChannelNewItem {
    return r._itemList
}
