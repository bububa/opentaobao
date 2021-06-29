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
type TaobaoPlaceStoreTagsUpdateRequest struct {
    model.Params
    // 门店信息
    _storeUpdate   *StoreUpdateTopDto
    // 新增标list
    _addTags   []int64
    // 删除标list
    _removeTags   []int64
}

// 初始化TaobaoPlaceStoreTagsUpdateRequest对象
func NewTaobaoPlaceStoreTagsUpdateRequest() *TaobaoPlaceStoreTagsUpdateRequest{
    return &TaobaoPlaceStoreTagsUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreTagsUpdateRequest) GetApiMethodName() string {
    return "taobao.place.store.tags.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreTagsUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreUpdate Setter
// 门店信息
func (r *TaobaoPlaceStoreTagsUpdateRequest) SetStoreUpdate(_storeUpdate *StoreUpdateTopDto) error {
    r._storeUpdate = _storeUpdate
    r.Set("store_update", _storeUpdate)
    return nil
}

// StoreUpdate Getter
func (r TaobaoPlaceStoreTagsUpdateRequest) GetStoreUpdate() *StoreUpdateTopDto {
    return r._storeUpdate
}
// AddTags Setter
// 新增标list
func (r *TaobaoPlaceStoreTagsUpdateRequest) SetAddTags(_addTags []int64) error {
    r._addTags = _addTags
    r.Set("add_tags", _addTags)
    return nil
}

// AddTags Getter
func (r TaobaoPlaceStoreTagsUpdateRequest) GetAddTags() []int64 {
    return r._addTags
}
// RemoveTags Setter
// 删除标list
func (r *TaobaoPlaceStoreTagsUpdateRequest) SetRemoveTags(_removeTags []int64) error {
    r._removeTags = _removeTags
    r.Set("remove_tags", _removeTags)
    return nil
}

// RemoveTags Getter
func (r TaobaoPlaceStoreTagsUpdateRequest) GetRemoveTags() []int64 {
    return r._removeTags
}
