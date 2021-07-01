package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店团购套餐关联适用门店 API请求
alitrip.tuan.hotel.adapt.store.get

输入shid，返回关联门店详情信息
*/
type AlitripTuanHotelAdaptStoreGetAPIRequest struct {
    model.Params
    // 标准酒店ID列表,逗号分割。与hid_list二者只能选一
    _shidList   []int64
    // 物理酒店ID列表，逗号分割。与shid_list二者只能选一
    _hidList   []int64
}

// 初始化AlitripTuanHotelAdaptStoreGetAPIRequest对象
func NewAlitripTuanHotelAdaptStoreGetRequest() *AlitripTuanHotelAdaptStoreGetAPIRequest{
    return &AlitripTuanHotelAdaptStoreGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.adapt.store.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShidList Setter
// 标准酒店ID列表,逗号分割。与hid_list二者只能选一
func (r *AlitripTuanHotelAdaptStoreGetAPIRequest) SetShidList(_shidList []int64) error {
    r._shidList = _shidList
    r.Set("shid_list", _shidList)
    return nil
}

// ShidList Getter
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetShidList() []int64 {
    return r._shidList
}
// HidList Setter
// 物理酒店ID列表，逗号分割。与shid_list二者只能选一
func (r *AlitripTuanHotelAdaptStoreGetAPIRequest) SetHidList(_hidList []int64) error {
    r._hidList = _hidList
    r.Set("hid_list", _hidList)
    return nil
}

// HidList Getter
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetHidList() []int64 {
    return r._hidList
}
