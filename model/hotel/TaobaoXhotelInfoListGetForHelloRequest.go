package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
哈罗获取酒店详情 API请求
taobao.xhotel.info.list.get.for.hello

哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
*/
type TaobaoXhotelInfoListGetForHelloRequest struct {
    model.Params
    // 参数封装模型
    _hotelInfoParam   *HotelInfoParam
}

// 初始化TaobaoXhotelInfoListGetForHelloRequest对象
func NewTaobaoXhotelInfoListGetForHelloRequest() *TaobaoXhotelInfoListGetForHelloRequest{
    return &TaobaoXhotelInfoListGetForHelloRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelInfoListGetForHelloRequest) GetApiMethodName() string {
    return "taobao.xhotel.info.list.get.for.hello"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelInfoListGetForHelloRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HotelInfoParam Setter
// 参数封装模型
func (r *TaobaoXhotelInfoListGetForHelloRequest) SetHotelInfoParam(_hotelInfoParam *HotelInfoParam) error {
    r._hotelInfoParam = _hotelInfoParam
    r.Set("hotel_info_param", _hotelInfoParam)
    return nil
}

// HotelInfoParam Getter
func (r TaobaoXhotelInfoListGetForHelloRequest) GetHotelInfoParam() *HotelInfoParam {
    return r._hotelInfoParam
}
