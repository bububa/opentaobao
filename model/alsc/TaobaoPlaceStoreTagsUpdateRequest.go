package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店打标去标 API请求
taobao.place.store.tags.update

门店打标去标
*/
type TaobaoPlaceStoreTagsUpdateAPIRequest struct {
    model.Params
    // 门店信息
    _storeUpdate   *StoreUpdateTopDTO
    // 新增标list
    _addTags   []int64
    // 删除标list
    _removeTags   []int64
}

// 初始化TaobaoPlaceStoreTagsUpdateAPIRequest对象
func NewTaobaoPlaceStoreTagsUpdateRequest() *TaobaoPlaceStoreTagsUpdateAPIRequest{
    return &TaobaoPlaceStoreTagsUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreTagsUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.place.store.tags.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreTagsUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreUpdate Setter
// 门店信息
func (r *TaobaoPlaceStoreTagsUpdateAPIRequest) SetStoreUpdate(_storeUpdate *StoreUpdateTopDTO) error {
    r._storeUpdate = _storeUpdate
    r.Set("store_update", _storeUpdate)
    return nil
}

// StoreUpdate Getter
func (r TaobaoPlaceStoreTagsUpdateAPIRequest) GetStoreUpdate() *StoreUpdateTopDTO {
    return r._storeUpdate
}
// AddTags Setter
// 新增标list
func (r *TaobaoPlaceStoreTagsUpdateAPIRequest) SetAddTags(_addTags []int64) error {
    r._addTags = _addTags
    r.Set("add_tags", _addTags)
    return nil
}

// AddTags Getter
func (r TaobaoPlaceStoreTagsUpdateAPIRequest) GetAddTags() []int64 {
    return r._addTags
}
// RemoveTags Setter
// 删除标list
func (r *TaobaoPlaceStoreTagsUpdateAPIRequest) SetRemoveTags(_removeTags []int64) error {
    r._removeTags = _removeTags
    r.Set("remove_tags", _removeTags)
    return nil
}

// RemoveTags Getter
func (r TaobaoPlaceStoreTagsUpdateAPIRequest) GetRemoveTags() []int64 {
    return r._removeTags
}
