package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫品牌新品同步API APIRequest
tmall.brand.item.upload

支撑天猫品牌将各渠道新品信息同步至平台
*/
type TmallBrandItemUploadRequest struct {
    model.Params

    // 需要同步的商品列表
    itemList   []TmallBrandChannelNewItem 

}

func NewTmallBrandItemUploadRequest() *TmallBrandItemUploadRequest{
    return &TmallBrandItemUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallBrandItemUploadRequest) GetApiMethodName() string {
    return "tmall.brand.item.upload"
}

func (r TmallBrandItemUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallBrandItemUploadRequest) SetItemList(itemList []TmallBrandChannelNewItem) error {
    r.itemList = itemList
    r.Set("item_list", itemList)
    return nil
}

func (r TmallBrandItemUploadRequest) GetItemList() []TmallBrandChannelNewItem {
    return r.itemList
}

