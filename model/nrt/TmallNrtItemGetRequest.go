package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售商品信息查询 API请求
tmall.nrt.item.get

查询新零售商品信息
*/
type TmallNrtItemGetRequest struct {
    model.Params
    // 城市站id
    boothId   int64
    // 商品id
    itemId   int64
}

// 初始化TmallNrtItemGetRequest对象
func NewTmallNrtItemGetRequest() *TmallNrtItemGetRequest{
    return &TmallNrtItemGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtItemGetRequest) GetApiMethodName() string {
    return "tmall.nrt.item.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BoothId Setter
// 城市站id
func (r *TmallNrtItemGetRequest) SetBoothId(boothId int64) error {
    r.boothId = boothId
    r.Set("booth_id", boothId)
    return nil
}

// BoothId Getter
func (r TmallNrtItemGetRequest) GetBoothId() int64 {
    return r.boothId
}
// ItemId Setter
// 商品id
func (r *TmallNrtItemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallNrtItemGetRequest) GetItemId() int64 {
    return r.itemId
}
