package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
详情页静态信息 
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口
*/
func AlitripHotelDetailStaticinfoGet(clt *core.SDKClient, req *hotel.AlitripHotelDetailStaticinfoGetRequest, session string) (*hotel.AlitripHotelDetailStaticinfoGetResponse, error) {
    var resp hotel.AlitripHotelDetailStaticinfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
