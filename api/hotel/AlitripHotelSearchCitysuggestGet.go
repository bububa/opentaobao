package hotel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/hotel"
)

/* 
城市Suggest接口 
alitrip.hotel.search.citysuggest.get

城市Suggest接口
*/
func AlitripHotelSearchCitysuggestGet(clt *core.SDKClient, req *hotel.AlitripHotelSearchCitysuggestGetAPIRequest, session string) (*hotel.AlitripHotelSearchCitysuggestGetAPIResponse, error) {
    var resp hotel.AlitripHotelSearchCitysuggestGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
