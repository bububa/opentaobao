package hotelalliance

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotelalliance"
)

/* 
获取单体酒店信息 
alitrip.hotel.single.info.get

用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务
*/
func AlitripHotelSingleInfoGet(clt *core.SDKClient, req *hotelalliance.AlitripHotelSingleInfoGetRequest, session string) (*hotelalliance.AlitripHotelSingleInfoGetAPIResponse, error) {
    var resp hotelalliance.AlitripHotelSingleInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
