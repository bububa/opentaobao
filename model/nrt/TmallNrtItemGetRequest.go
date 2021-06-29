package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售商品信息查询 APIRequest
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

func NewTmallNrtItemGetRequest() *TmallNrtItemGetRequest{
    return &TmallNrtItemGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtItemGetRequest) GetApiMethodName() string {
    return "tmall.nrt.item.get"
}

func (r TmallNrtItemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtItemGetRequest) SetBoothId(boothId int64) error {
    r.boothId = boothId
    r.Set("booth_id", boothId)
    return nil
}

func (r TmallNrtItemGetRequest) GetBoothId() int64 {
    return r.boothId
}

func (r *TmallNrtItemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallNrtItemGetRequest) GetItemId() int64 {
    return r.itemId
}

