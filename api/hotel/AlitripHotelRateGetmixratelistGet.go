package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
酒店评论接口 
alitrip.hotel.rate.getmixratelist.get

酒店评论接口
*/
func AlitripHotelRateGetmixratelistGet(clt *core.SDKClient, req *hotel.AlitripHotelRateGetmixratelistGetRequest, session string) (*hotel.AlitripHotelRateGetmixratelistGetResponse, error) {
    var resp hotel.AlitripHotelRateGetmixratelistGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
