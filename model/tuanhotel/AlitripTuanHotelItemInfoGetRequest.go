package tuanhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
宝贝信息查询接口 API请求
alitrip.tuan.hotel.item.info.get

商家查询发布的宝贝详情信息
*/
type AlitripTuanHotelItemInfoGetRequest struct {
    model.Params
    // 宝贝ID
    _itemId   int64
}

// 初始化AlitripTuanHotelItemInfoGetRequest对象
func NewAlitripTuanHotelItemInfoGetRequest() *AlitripTuanHotelItemInfoGetRequest{
    return &AlitripTuanHotelItemInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelItemInfoGetRequest) GetApiMethodName() string {
    return "alitrip.tuan.hotel.item.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelItemInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 宝贝ID
func (r *AlitripTuanHotelItemInfoGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlitripTuanHotelItemInfoGetRequest) GetItemId() int64 {
    return r._itemId
}
