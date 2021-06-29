package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店团购套餐关联适用门店 APIRequest
alitrip.tuan.hotel.adapt.store.get

输入shid，返回关联门店详情信息
*/
type AlitripTuanHotelAdaptStoreGetRequest struct {
    model.Params

    // 标准酒店ID列表,逗号分割。与hid_list二者只能选一
    shidList   []int64 

    // 物理酒店ID列表，逗号分割。与shid_list二者只能选一
    hidList   []int64 

}

func NewAlitripTuanHotelAdaptStoreGetRequest() *AlitripTuanHotelAdaptStoreGetRequest{
    return &AlitripTuanHotelAdaptStoreGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTuanHotelAdaptStoreGetRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.adapt.store.get"
}

func (r AlitripTuanHotelAdaptStoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTuanHotelAdaptStoreGetRequest) SetShidList(shidList []int64) error {
    r.shidList = shidList
    r.Set("shid_list", shidList)
    return nil
}

func (r AlitripTuanHotelAdaptStoreGetRequest) GetShidList() []int64 {
    return r.shidList
}

func (r *AlitripTuanHotelAdaptStoreGetRequest) SetHidList(hidList []int64) error {
    r.hidList = hidList
    r.Set("hid_list", hidList)
    return nil
}

func (r AlitripTuanHotelAdaptStoreGetRequest) GetHidList() []int64 {
    return r.hidList
}

