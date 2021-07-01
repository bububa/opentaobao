package tuanhotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tuanhotel"
)

/* 
宝贝信息查询接口 
alitrip.tuan.hotel.item.info.get

商家查询发布的宝贝详情信息
*/
func AlitripTuanHotelItemInfoGet(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelItemInfoGetAPIRequest, session string) (*tuanhotel.AlitripTuanHotelItemInfoGetAPIResponse, error) {
    var resp tuanhotel.AlitripTuanHotelItemInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
