package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店打标去标 APIRequest
taobao.place.store.tags.update

门店打标去标
*/
type TaobaoPlaceStoreTagsUpdateRequest struct {
    model.Params

    // 门店信息
    storeUpdate   *StoreUpdateTopDto 

    // 新增标list
    addTags   []Number 

    // 删除标list
    removeTags   []Number 

}

func NewTaobaoPlaceStoreTagsUpdateRequest() *TaobaoPlaceStoreTagsUpdateRequest{
    return &TaobaoPlaceStoreTagsUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoreTagsUpdateRequest) GetApiMethodName() string {
    return "taobao.place.store.tags.update"
}

func (r TaobaoPlaceStoreTagsUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoreTagsUpdateRequest) SetStoreUpdate(storeUpdate *StoreUpdateTopDto) error {
    r.storeUpdate = storeUpdate
    r.Set("store_update", storeUpdate)
    return nil
}

func (r TaobaoPlaceStoreTagsUpdateRequest) GetStoreUpdate() *StoreUpdateTopDto {
    return r.storeUpdate
}

func (r *TaobaoPlaceStoreTagsUpdateRequest) SetAddTags(addTags []Number) error {
    r.addTags = addTags
    r.Set("add_tags", addTags)
    return nil
}

func (r TaobaoPlaceStoreTagsUpdateRequest) GetAddTags() []Number {
    return r.addTags
}

func (r *TaobaoPlaceStoreTagsUpdateRequest) SetRemoveTags(removeTags []Number) error {
    r.removeTags = removeTags
    r.Set("remove_tags", removeTags)
    return nil
}

func (r TaobaoPlaceStoreTagsUpdateRequest) GetRemoveTags() []Number {
    return r.removeTags
}

