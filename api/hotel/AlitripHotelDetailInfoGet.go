package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
详情页动态信息接口 
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法
*/
func AlitripHotelDetailInfoGet(clt *core.SDKClient, req *hotel.AlitripHotelDetailInfoGetAPIRequest, session string) (*hotel.AlitripHotelDetailInfoGetAPIResponse, error) {
    var resp hotel.AlitripHotelDetailInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
