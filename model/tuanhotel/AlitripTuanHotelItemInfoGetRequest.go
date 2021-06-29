package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
宝贝信息查询接口 APIRequest
alitrip.tuan.hotel.item.info.get

商家查询发布的宝贝详情信息
*/
type AlitripTuanHotelItemInfoGetRequest struct {
    model.Params

    // 宝贝ID
    itemId   int64 

}

func NewAlitripTuanHotelItemInfoGetRequest() *AlitripTuanHotelItemInfoGetRequest{
    return &AlitripTuanHotelItemInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTuanHotelItemInfoGetRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.item.info.get"
}

func (r AlitripTuanHotelItemInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTuanHotelItemInfoGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlitripTuanHotelItemInfoGetRequest) GetItemId() int64 {
    return r.itemId
}

