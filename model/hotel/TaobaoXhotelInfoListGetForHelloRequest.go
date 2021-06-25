package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
哈罗获取酒店详情 APIRequest
taobao.xhotel.info.list.get.for.hello

哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
*/
type TaobaoXhotelInfoListGetForHelloRequest struct {
    model.Params

    // 参数封装模型
    hotelInfoParam   *HotelInfoParam 

}

func NewTaobaoXhotelInfoListGetForHelloRequest() *TaobaoXhotelInfoListGetForHelloRequest{
    return &TaobaoXhotelInfoListGetForHelloRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelInfoListGetForHelloRequest) GetApiMethodName() string {
    return "taobao.xhotel.info.list.get.for.hello"
}

func (r TaobaoXhotelInfoListGetForHelloRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelInfoListGetForHelloRequest) SetHotelInfoParam(hotelInfoParam *HotelInfoParam) error {
    r.hotelInfoParam = hotelInfoParam
    r.Set("hotel_info_param", hotelInfoParam)
    return nil
}

func (r TaobaoXhotelInfoListGetForHelloRequest) GetHotelInfoParam() *HotelInfoParam {
    return r.hotelInfoParam
}

