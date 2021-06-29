package hotelalliance

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotelalliance"
)

/* 
获取合作商信息 
alitrip.hotel.hms.partner.info.get

用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
*/
func AlitripHotelHmsPartnerInfoGet(clt *core.SDKClient, req *hotelalliance.AlitripHotelHmsPartnerInfoGetRequest, session string) (*hotelalliance.AlitripHotelHmsPartnerInfoGetAPIResponse, error) {
    var resp hotelalliance.AlitripHotelHmsPartnerInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
