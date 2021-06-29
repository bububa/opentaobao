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
    storeUpdate   *StoreUpdateTopDto
    // 新增标list
    addTags   []int64
    // 删除标list
    removeTags   []int64
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
func (r *TaobaoPlaceStoreTagsUpdateRequest) SetStoreUpdate(storeUpdate *StoreUpdateTopDto) error {
    r.storeUpdate = storeUpdate
    r.Set("store_update", storeUpdate)
    return nil
}

// StoreUpdate Getter
func (r TaobaoPlaceStoreTagsUpdateRequest) GetStoreUpdate() *StoreUpdateTopDto {
    return r.storeUpdate
}
// AddTags Setter
// 新增标list
func (r *TaobaoPlaceStoreTagsUpdateRequest) SetAddTags(addTags []int64) error {
    r.addTags = addTags
    r.Set("add_tags", addTags)
    return nil
}

// AddTags Getter
func (r TaobaoPlaceStoreTagsUpdateRequest) GetAddTags() []int64 {
    return r.addTags
}
// RemoveTags Setter
// 删除标list
func (r *TaobaoPlaceStoreTagsUpdateRequest) SetRemoveTags(removeTags []int64) error {
    r.removeTags = removeTags
    r.Set("remove_tags", removeTags)
    return nil
}

// RemoveTags Getter
func (r TaobaoPlaceStoreTagsUpdateRequest) GetRemoveTags() []int64 {
    return r.removeTags
}
